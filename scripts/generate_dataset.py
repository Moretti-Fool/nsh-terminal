import json
import os
import random
from openrouter_adapter import OpenRouterAdapter

def generate_batch(adapter, domain, shell, os_name, count=10):
    sys_prompt = "You are a master dataset generator for a Natural Language to CLI tool (nsh). You must return ONLY raw JSON, without Markdown blocks."
    
    user_prompt = r'''
Generate exactly {count} diverse, highly realistic training examples for translating natural language into {shell} commands on {os_name}.
Focus heavily on the domain: {domain}. Include a mix of BASIC (easy) and ADVANCED (complex) commands.

Make the natural language queries realistic (some polite, some shorthand, some misspelled, like 'how do i restart wsl', 'rm all node_modules', 'find port 3000').

Return a JSON array of objects, where each object has:
- "nl": The natural language query from the user.
- "explanation": A very brief conversational explanation of what the command will do (1 sentence).
- "commands": A list of string commands to execute (valid {shell}).

Example Output format exactly like this:
[
  {
    "nl": "find all text files recursively",
    "explanation": "I will search the current directory and all subdirectories for text files.",
    "commands": ["command_here depending on shell"]
  }
]
'''
    user_prompt = user_prompt.replace("{count}", str(count)).replace("{domain}", domain).replace("{shell}", shell).replace("{os_name}", os_name)
    
    try:
        response = adapter.ask(user_prompt, sys_prompt)
        resp = response.strip()
        if resp.startswith("`json"):
            resp = resp.replace("`json", "", 1)
        if resp.endswith("`"):
            resp = resp[:-3]
        resp = resp.strip()
        
        data = json.loads(resp)
        return data
    except Exception as e:
        print(f"Failed generating {domain} for {shell}: {e}")
        return []

def main():
    adapter = OpenRouterAdapter("../.env")
    
    domains = [
        "Git version control (branching, merging, rebasing, logs)",
        "Docker (containers, images, compose, pruning)",
        "Networking (ports, IPs, DNS, testing connections)",
        "Filesystem (finding files by date/size, renaming, moving, reading text)",
        "Process Management (killing tasks, checking CPU usage, listing services)",
        "Package Managers (npm, pip, choco, apt, brew)",
        "System Info (disk space, OS version, memory, environment variables)"
    ]

    platforms = [
        {"os": "windows", "shell": "powershell"},
        {"os": "windows", "shell": "cmd"},
        {"os": "linux", "shell": "bash"}
    ]

    out_file = "../training/data/nsh_train_advanced.jsonl"
    
    total_added = 0
    with open(out_file, 'w', encoding='utf-8') as f:
        # For each platform, generate 1 batch per domain (3 platforms * 7 domains * 15 samples = ~315 examples)
        for platform in platforms:
            os_name = platform["os"]
            shell = platform["shell"]
            
            system_message = f"You are nsh, a terminal command translator.\nOS: {os_name}\nShell: {shell}\nCWD: /home/user/workspace\n\nWhen you are done investigating and want to give the final commands to the user, respond with JSON ONLY:\n{{\"explanation\": \"any conversational text you want to say to the user\", \"commands\":[\"...\"],\"dialect\":\"{shell}\"}}\ndialect must be exactly {shell}.\nAt most 3 commands. No markdown.\nDo not invent filenames that are not in the listing or tool results."
            
            for i, domain in enumerate(domains):
                print(f"Generating batch {i+1}/{len(domains)}: {domain} for {shell} on {os_name}...")
                batch = generate_batch(adapter, domain, shell, os_name, count=15)
                for item in batch:
                    if 'nl' not in item or 'commands' not in item or 'explanation' not in item:
                        continue
                    
                    ast = {
                        "explanation": item["explanation"],
                        "commands": item["commands"],
                        "dialect": shell
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
                
    print(f"\nDone! Generated {total_added} diverse shell examples in {out_file}.")

if __name__ == '__main__':
    main()
