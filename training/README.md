# nsh Model Fine-tuning

This directory contains the QLoRA fine-tuning script for nsh's natural language to shell command translation model. The script is optimized to run on an RTX 3050 Ti (4GB VRAM) using Unsloth.

## Setup Instructions

1. Ensure you have Python 3.10+ and CUDA 12.7 installed.
2. Create and activate a virtual environment:
   ```bash
   python -m venv venv
   # On Windows:
   .\venv\Scripts\activate
   # On Linux:
   source venv/bin/activate
   ```
3. Install dependencies:
   ```bash
   pip install -r requirements.txt
   ```

## Running Training

Run the training script from within this `training` directory:
```bash
python train.py
```
**Expected Training Time:** Approximately 2-4 hours on an RTX 3050 Ti.

The script will save:
- Checkpoints to `output/checkpoints`
- The LoRA adapter to `output/nsh-qwen-lora`
- The fully merged model to `output/nsh-qwen-merged`

## Export to GGUF

To use the model with Ollama, you'll need to export the merged model to GGUF format. Using `llama.cpp`:

```bash
git clone https://github.com/ggerganov/llama.cpp
cd llama.cpp
pip install -r requirements.txt
python convert_hf_to_gguf.py ../output/nsh-qwen-merged --outfile ../output/nsh-qwen-merged.gguf --outtype q4_k_m
```

## Registering with Ollama

Create a `Modelfile`:
```text
FROM ./output/nsh-qwen-merged.gguf
TEMPLATE """{{ if .System }}<|im_start|>system
{{ .System }}<|im_end|>
{{ end }}{{ if .Prompt }}<|im_start|>user
{{ .Prompt }}<|im_end|>
{{ end }}<|im_start|>assistant
"""
```

Then create and run the model:
```bash
ollama create nsh-qwen -f Modelfile
ollama run nsh-qwen
```

## Troubleshooting

### Common VRAM OOM Errors
If you run into Out of Memory (OOM) errors on your 4GB GPU:
1. **Reduce Batch Size:** Keep `per_device_train_batch_size=1`.
2. **Reduce Max Sequence Length:** In `train.py`, reduce `max_seq_length` from 512 to 256 or 128 if your inputs are short.
3. **Close other apps:** Ensure browsers, Discord, and other hardware-accelerated apps are closed during training.
4. **Gradient Checkpointing:** Ensure `use_gradient_checkpointing="unsloth"` is active (it is by default in this script).
