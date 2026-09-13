import json
import os
import requests
import random
import time

OLLAMA_URL = "http://localhost:11434/api/generate"
MODEL = "qwen2.5:7b-instruct-q4_K_M"

def generate_powershell_pairs(count=20):
    prompt = f"""You are a helpful assistant generating synthetic training data for a PowerShell terminal translator.
Generate exactly {count} distinct user requests and their corresponding Windows PowerShell commands.
Vary the difficulty. Include file management, networking, system info, and services.
Format your response as a STRICT JSON array of objects, like this:
[
  {{"query": "find all python files", "command": "Get-ChildItem -Filter *.py -Recurse"}},
  {{"query": "restart the spooler service", "command": "Restart-Service -Name Spooler"}}
]
Do not output anything else. Just the JSON array.
"""
    try:
        print("Generating synthetic data via Ollama (this takes 1-2 mins)...")
        r = requests.post(OLLAMA_URL, json={
            "model": MODEL,
            "prompt": prompt,
            "format": "json",
            "stream": False
        }, timeout=180)
        resp = r.json()
        data = json.loads(resp["response"])
        return data
    except Exception as e:
        print(f"Failed to generate synthetic data: {e}")
        return []

def append_to_dataset(pairs):
    if not pairs:
        return

    from prepare_agentic_dataset import make_system_prompt
    
    dataset = []
    
    for pair in pairs:
        query = pair.get("query", "")
        cmd = pair.get("command", "")
        if not query or not cmd:
            continue
            
        messages = [
            {"role": "system", "content": make_system_prompt("windows", "powershell")},
            {"role": "user", "content": query},
            {"role": "assistant", "content": json.dumps({"commands": [cmd], "dialect": "powershell"})}
        ]
        dataset.append({"messages": messages})
        
    # Append to existing train file
    with open("data/nsh_train.jsonl", "a", encoding="utf-8") as f:
        for d in dataset:
            f.write(json.dumps(d) + "\n")
            
    print(f"Appended {len(dataset)} new synthetic examples to the training dataset!")

if __name__ == "__main__":
    pairs = generate_powershell_pairs(20)
    append_to_dataset(pairs)
