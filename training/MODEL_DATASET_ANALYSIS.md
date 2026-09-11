# Natural Language to Shell Translation: Model & Dataset Analysis

## Executive Summary

During testing of `nsh` on Windows using the locally fine-tuned `nsh-local` model, issuing natural language commands such as `check disk` produced unexpected commands (`Get-ComputerInfo`) instead of appropriate disk utilities (`Get-Volume`, `Get-PSDrive`, or `chkdsk C:`).

This document details the root causes identified across the training pipeline, reviews open-source datasets available on the internet, and specifies an end-to-end solution for dataset curation, prompt alignment, and fine-tuning.

---

## Part 1: Root Cause Analysis of Current Behavior

### 1. Dialect & Shell Isolation Mismatch
In [`training/generate_dataset.py`](./generate_dataset.py), the query `"check disk"` was defined **only** under `cmd_templates`:
```python
# training/generate_dataset.py lines 597-601
{
    "queries": ["check disk", "chkdsk"],
    "commands": ["chkdsk C:"],
    "dialect": "cmd"
}
```
At runtime on Windows, `nsh` defaults to `powershell`. The prompt sent to Ollama explicitly enforces:
```text
Shell: powershell
Respond with JSON only: {"commands":["..."],"dialect":"powershell"}
dialect must be exactly powershell.
```
Because the model was instructed to output dialect `powershell`, it could not emit the `cmd` training sample (`chkdsk C:`).

### 2. Missing PowerShell Command Coverage
In `powershell_templates`, there were **zero** disk-related commands (`Get-Volume`, `Get-PSDrive`, `Get-Disk`, or calling `chkdsk`).
However, `powershell_templates` did define:
```python
# training/generate_dataset.py lines 70-78
{
    "queries": [
        "show system info",
        "get computer info",
        "what are my system specs"
    ],
    "commands": ["Get-ComputerInfo"],
    "dialect": "powershell"
}
```
Faced with an unfamiliar PowerShell query for `"check disk"`, the model fell back to the nearest system inspection cmdlet it was trained on: `Get-ComputerInfo`.

### 3. Synthetic Data Generation Flaws (Nonsensical OS/Shell Pairings)
In `training/generate_dataset.py`, line 689 generated prompts by independently choosing the OS and dialect:
```python
sys_prompt = f"You are a helpful assistant that translates natural language into shell commands. The current operating system is {random.choice(oss)}, the shell is {dialect}, and the current working directory is {cwd}..."
```
This generated thousands of impossible pairings in `nsh_train.jsonl` and `nsh_eval.jsonl`:
- `The current operating system is darwin, the shell is cmd...` *(Windows CMD on macOS)*
- `The current operating system is linux, the shell is cmd...` *(Windows CMD on Linux)*

This noise weakened the model's ability to ground specific commands to their true operating environments.

### 4. Prompt Template Divergence (Training vs. Runtime)
Fine-tuned small-parameter models (such as 7B or 1.5B/3B) are sensitive to system prompt formatting:
- **Fine-tuning Prompt**:
  `"You are a helpful assistant that translates natural language into shell commands. The current operating system is {os}, the shell is {dialect}..."`
- **`nsh` Runtime Prompt (`internal/ollama/chat.go`)**:
  `"You are nsh, a terminal command translator.\nOS: windows\nShell: powershell\nCWD: ...\nRespond with JSON only: ...\ndialect must be exactly ..."`

Because the runtime prompt differed in token structure, syntax, and instructions, the model struggled to activate the fine-tuned representations reliably.

---

## Part 2: Open-Source Datasets Available on the Internet

The following public datasets on Hugging Face and GitHub provide real-world, high-quality query-to-command mappings:

### 1. `carosh/cli-1m` (Recommended)
- **Hugging Face**: [carosh/cli-1m](https://huggingface.co/datasets/carosh/cli-1m)
- **Scale**: ~975,000+ pairs (Apache 2.0 license).
- **Shell Support**: Specifically covers **PowerShell**, **Bash**, **Zsh**, **Fish**, **Nu**, and **Oils-OSH**.
- **Domains**: 18 real-world domains (cloud, devops, system administration, networking, disk and file operations).
- **Curated Subsets**:
  - `sample`: Stratified 50,000-row subset.
  - `validation`: ~15,000 hand-curated, LIMA-quality pairs.
- **Python Usage**:
  ```python
  from datasets import load_dataset
  ds = load_dataset("carosh/cli-1m", name="sample")
  ```

### 2. `tldr-pages` / `neulab/tldr`
- **Hugging Face**: [neulab/tldr](https://huggingface.co/datasets/neulab/tldr)
- **GitHub**: [tldr-pages/tldr](https://github.com/tldr-pages/tldr)
- **Description**: Human-written command examples with clear natural language descriptions for thousands of CLI tools across Linux, macOS, and Windows.
- **Signal-to-Noise**: High signal, zero synthetic artifacts.

### 3. `westenfelder/NL2SH-ALFA`
- **Hugging Face**: [westenfelder/NL2SH-ALFA](https://huggingface.co/datasets/westenfelder/NL2SH-ALFA)
- **Description**: Unified, deduplicated, and filtered collection combining `NL2Bash`, `NL2CMD`, `tldr-pages`, and `InterCode-Bash`.

### 4. `AnishJoshi/nl2bash-custom`
- **Hugging Face**: [AnishJoshi/nl2bash-custom](https://huggingface.co/datasets/AnishJoshi/nl2bash-custom)
- **Description**: Formatted instruction-tuning dataset built from the classic `NL2Bash` benchmark.

---

## Part 3: Solution Architecture

### 1. Strict OS-to-Dialect Mapping
Ensure no synthetic generator pairs incompatible operating systems and shells:
- `windows` &rarr; `powershell` (default), `cmd`
- `linux` &rarr; `bash`, `sh`
- `darwin` &rarr; `zsh` (default), `bash`

### 2. Comprehensive Multi-Shell Task Matrix
Every core operational domain must have representative commands in all relevant dialects:

| Task | PowerShell (`windows`) | CMD (`windows`) | Bash / Zsh (`linux`, `darwin`) |
| :--- | :--- | :--- | :--- |
| **Check disk status/space** | `Get-Volume`, `Get-PSDrive -PSProvider FileSystem` | `chkdsk C:`, `wmic logicaldisk get ...` | `df -h` |
| **List directory** | `Get-ChildItem` | `dir` | `ls -la` |
| **Active network connections** | `Get-NetTCPConnection` | `netstat -ano` | `ss -tuln`, `netstat -tlpn` |
| **Resolve DNS** | `Resolve-DnsName -Name <domain>` | `nslookup <domain>` | `dig <domain>`, `nslookup <domain>` |
| **Find file by name** | `Get-ChildItem -Filter *.txt -Recurse` | `dir /s /b *.txt` | `find . -name "*.txt"` |
| **Find text in file** | `Select-String -Pattern "err" -Path ...` | `findstr "err" ...` | `grep -rn "err" ...` |
| **Process management** | `Get-Process`, `Stop-Process -Id ...` | `tasklist`, `taskkill /PID ... /F` | `ps aux`, `kill -9 ...` |
| **Service management** | `Get-Service`, `Restart-Service ...` | `net start`, `sc query` | `systemctl status ...` |

### 3. Exact Prompt Alignment with Runtime `nsh`
The training records in `nsh_train.jsonl` must match the exact prompt structure generated by `internal/ollama/chat.go`:
```json
{
  "messages": [
    {
      "role": "system",
      "content": "You are nsh, a terminal command translator.\nOS: windows\nShell: powershell\nCWD: C:\\Users\\User\n\nRespond with JSON only: {\"commands\":[\"...\"],\"dialect\":\"powershell\"}\ndialect must be exactly powershell.\nAt most 3 commands. No markdown. No explanations.\nDo not invent filenames that are not in the listing or tool results."
    },
    {
      "role": "user",
      "content": "check disk"
    },
    {
      "role": "assistant",
      "content": "{\"commands\":[\"Get-Volume\"],\"dialect\":\"powershell\"}"
    }
  ]
}
```

---

## Part 4: Implementation Roadmap

1. **Dataset Pipeline Script** (`training/prepare_real_dataset.py`):
   - Ingest data from `carosh/cli-1m` and `tldr-pages`.
   - Filter for `powershell`, `cmd`, `bash`, and `zsh`.
   - Re-encode into the exact `nsh` system prompt format.
2. **Quality Curation**:
   - Produce a balanced dataset of 10,000–20,000 high-quality samples.
   - 40% PowerShell / Windows CMD, 40% Linux Bash, 20% macOS Zsh.
3. **Fine-Tuning**:
   - Use QLoRA with Unsloth on `Qwen2.5-Coder-7B-Instruct` or smaller base models.
   - Export directly to GGUF for local Ollama serving (`ollama create nsh-local -f Modelfile`).
4. **Validation**:
   - Evaluate against `nsh_eval.jsonl` ensuring zero syntax errors and valid shell commands across all platforms.

---

## Part 5: Small Pre-Trained & Fine-Tuned Models on the Internet

### 1. Dedicated NL-to-Shell Fine-Tuned Models
*   **`rlawltjd/code-llama3-7b-text-to-bash`** (Hugging Face): Fine-tuned CodeLlama on data-augmented natural language to bash tasks.
*   **`Edoigtrd/T5-nl2bash`** (Hugging Face): Lightweight T5-based model (0.2B parameters) specialized strictly for NL2Bash translation.
*   **`laion/rl_r2egym-nl2bash-swesmith`** (Hugging Face): Reinforcement-learning tuned model on bash generation benchmarks.
*   *Limitation of existing public fine-tunes*: The overwhelming majority of public fine-tuned models are **Linux Bash only**. Almost none of them natively know PowerShell or Windows CMD idioms.

### 2. Recommended Small Base Models for Multi-Shell Fine-Tuning
For local serving in `nsh` via Ollama, modern small coder architectures offer high native knowledge of both Bash and PowerShell:

*   **`Qwen2.5-Coder-1.5B-Instruct`** (`ollama run qwen2.5-coder:1.5b`):
    - **Size**: ~1.5B parameters (~1.0 GB in Q4_K_M).
    - **Characteristics**: Extremely fast inference with sub-100ms cold starts. Ideal for low-latency command suggestions and resource-constrained environments.
*   **`Qwen2.5-Coder-3B-Instruct`** (`ollama run qwen2.5-coder:3b`):
    - **Size**: ~3B parameters (~1.9 GB in Q4_K_M).
    - **Characteristics**: Optimal balance between parameter efficiency and deep syntactic reasoning. Accurately handles multi-parameter PowerShell cmdlets and complex piped shell commands.
*   **`ibm-granite/granite-3b-code-instruct`**:
    - **Size**: 3B parameters. Specifically benchmarked on natural language to Bash and PowerShell execution.
*   **`Qwen2.5-Coder-7B-Instruct`** (`ollama run qwen2.5-coder:7b`):
    - **Size**: ~7B parameters (~4.7 GB in Q4_K_M).
    - **Characteristics**: State-of-the-art capability for local coding models. Highly recommended when paired with curated `cli-1m` SFT.
