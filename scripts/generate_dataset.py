import json
import os
import random
from openrouter_adapter import OpenRouterAdapter

def generate_batch(adapter, domain, count=10):
    sys_prompt = "You are a master dataset generator for a Natural Language to CLI tool (nsh). You must return ONLY raw JSON, without Markdown blocks."
    
    user_prompt = r'''
Generate exactly {count} diverse, highly realistic training examples for translating natural language into Windows PowerShell commands.
Focus heavily on the domain: {domain}.

Make the natural language queries realistic (some polite, some shorthand, some misspelled, like 'how do i restart wsl', 'rm all node_modules', 'find port 3000').

Return a JSON array of objects, where each object has:
- "nl": The natural language query from the user.
- "explanation": A very brief conversational explanation of what the command will do (1 sentence).
- "commands": A list of string commands to execute (valid PowerShell).

Example Output format exactly like this:
[
  {
    "nl": "find all text files recursively",
    "explanation": "I will search the current directory and all subdirectories for text files.",
    "commands": ["Get-ChildItem -Filter *.txt -Recurse"]
  },
  {
    "nl": "kill process listening on port 8080",
    "explanation": "I will find the process ID using port 8080 and terminate it.",
    "commands": ["$pid = (Get-NetTCPConnection -LocalPort 8080).OwningProcess", "Stop-Process -Id $pid -Force"]
  }
]
'''
    user_prompt = user_prompt.replace("{count}", str(count)).replace("{domain}", domain)
    
    try:
        response = adapter.ask(user_prompt, sys_prompt)
        # basic sanitization
        resp = response.strip()
        if resp.startswith("`json"):
            resp = resp.replace("`json", "", 1)
        if resp.endswith("`"):
            resp = resp[:-3]
        resp = resp.strip()
        
        data = json.loads(resp)
        return data
    except Exception as e:
        print(f"Failed generating {domain}: {e}")
        return []

def main():
    adapter = OpenRouterAdapter("../.env")
    
    domains = [
        "Git version control (branching, merging, rebasing, logs)",
        "Docker (containers, images, compose, pruning)",
        "Networking (ports, IPs, DNS, testing connections)",
        "Filesystem (finding files by date/size, renaming, moving, reading text)",
        "Process Management (killing tasks, checking CPU usage, listing services)",
        "Package Managers (npm, pip, choco, winget)",
        "System Info (disk space, OS version, memory, environment variables)"
    ] * 2  # 14 batches of 20 = 280 examples

    out_file = "../training/data/nsh_train_advanced.jsonl"
    
    system_message = "You are nsh, a terminal command translator.\nOS: windows\nShell: powershell\nCWD: C:\\Users\\User\\Documents\n\nWhen you are done investigating and want to give the final commands to the user, respond with JSON ONLY:\n{\"explanation\": \"any conversational text you want to say to the user\", \"commands\":[\"...\"],\"dialect\":\"powershell\"}\ndialect must be exactly powershell.\nAt most 3 commands. No markdown.\nDo not invent filenames that are not in the listing or tool results."

    total_added = 0
    with open(out_file, 'w', encoding='utf-8') as f:
        for i, domain in enumerate(domains):
            print(f"Generating batch {i+1}/{len(domains)}: {domain}...")
            batch = generate_batch(adapter, domain, count=20)
            for item in batch:
                if 'nl' not in item or 'commands' not in item or 'explanation' not in item:
                    continue
                
                ast = {
                    "explanation": item["explanation"],
                    "commands": item["commands"],
                    "dialect": "powershell"
                }
                
                row = {
                    "messages": [
                        {"role": "system", "content": system_message},
                        {"role": "user", "content": item["nl"]},
                        {"role": "assistant", "content": json.dumps(ast)}
                    ]
                }
                f.write(json.dumps(row) + "\n")
                total_added += 1
                
    print(f"\nDone! Generated {total_added} advanced examples in {out_file}.")

if __name__ == '__main__':
    main()
