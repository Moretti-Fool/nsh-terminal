import json
import random
import os

random.seed(42)

# Ensure data directory exists
os.makedirs("data", exist_ok=True)

# Define the precise system prompt the shell uses at runtime
def make_system_prompt(os_name="windows", shell="powershell"):
    cwd = "C:\\Users\\User\\Documents" if os_name == "windows" else "/home/user/docs"
    return f"""You are nsh, a terminal command translator.
OS: {os_name}
Shell: {shell}
CWD: {cwd}

Respond with JSON only: {{"commands":["..."],"dialect":"{shell}"}}
dialect must be exactly {shell}.
At most 3 commands. No markdown. No explanations.
You have tools to investigate this machine before generating commands.
Use run_command to execute shell commands and read their output when you need runtime data (ports, processes, disk info, etc.).
Investigate first, then emit the JSON plan. You may call tools multiple times."""

# We will generate conversations
dataset = []

def add_example(user_query, tool_calls, final_commands, os_name="windows", shell="powershell"):
    messages = [
        {"role": "system", "content": make_system_prompt(os_name, shell)},
        {"role": "user", "content": user_query}
    ]
    
    # Add tool calls (investigation phase)
    for tc in tool_calls:
        # The assistant emits the tool call
        messages.append({
            "role": "assistant",
            "tool_calls": [
                {
                    "type": "function",
                    "function": {
                        "name": "run_command",
                        "arguments": json.dumps({"command": tc["cmd"]})
                    }
                }
            ]
        })
        # The tool responds
        messages.append({
            "role": "tool",
            "name": "run_command",
            "content": tc["out"]
        })

    # Finally, emit the correct JSON plan
    plan = {
        "commands": final_commands,
        "dialect": shell
    }
    messages.append({
        "role": "assistant",
        "content": json.dumps(plan, separators=(',', ':'))
    })
    
    dataset.append({"messages": messages})

# 1. Disk / System checks (Fixing the Get-ComputerInfo bug)
for i in range(50):
    add_example(
        "check disk space",
        [],
        ["Get-Volume"],
        "windows", "powershell"
    )
    add_example(
        "how much space is left on C drive?",
        [],
        ["Get-PSDrive -PSProvider FileSystem -Name C"],
        "windows", "powershell"
    )
    add_example(
        "check disk",
        [],
        ["Get-Volume"],
        "windows", "powershell"
    )

# 2. Port and Process checks using tool investigation
for i in range(50):
    # Port checking
    add_example(
        "what is running on port 8000?",
        [
            {"cmd": "Get-NetTCPConnection -LocalPort 8000 -ErrorAction SilentlyContinue", "out": "LocalAddress LocalPort OwningProcess\n------------ --------- -------------\n0.0.0.0      8000      14532"},
            {"cmd": "Get-Process -Id 14532 -ErrorAction SilentlyContinue", "out": "Handles  NPM(K)    PM(K)      WS(K)     CPU(s)     Id  SI ProcessName\n-------  ------    -----      -----     ------     --  -- -----------\n    150      12    12000      15000       1.23  14532   1 node"}
        ],
        ["Get-NetTCPConnection -LocalPort 8000 | Select-Object -Property LocalAddress,LocalPort,OwningProcess"],
        "windows", "powershell"
    )
    
    # Process investigation before killing
    add_example(
        "kill the process on port 8080",
        [
            {"cmd": "Get-NetTCPConnection -LocalPort 8080 -ErrorAction SilentlyContinue", "out": "LocalAddress LocalPort OwningProcess\n------------ --------- -------------\n0.0.0.0      8080      8833"}
        ],
        ["Stop-Process -Id 8833 -Force"],
        "windows", "powershell"
    )

# 3. Add repair examples (the model handles error feedback)
for i in range(50):
    msg_repair = [
        {"role": "system", "content": make_system_prompt("windows", "powershell")},
        {"role": "user", "content": "check disk"},
        {"role": "assistant", "content": '{"commands":["Get-ComputerInfo"],"dialect":"powershell"}'},
        {"role": "user", "content": "The previous command failed with this error:\nGet-ComputerInfo : The term 'Get-ComputerInfo' is not recognized...\nEmit a corrected JSON plan for the original request. Same OS and shell."},
        {"role": "assistant", "content": '{"commands":["Get-Volume"],"dialect":"powershell"}'}
    ]
    dataset.append({"messages": msg_repair})
    
    msg_repair_port = [
        {"role": "system", "content": make_system_prompt("windows", "powershell")},
        {"role": "user", "content": "show port 3000"},
        {"role": "assistant", "content": '{"commands":["Get-NetTCPConnection -Port 3000"],"dialect":"powershell"}'},
        {"role": "user", "content": "The previous command failed with this error:\nA parameter cannot be found that matches parameter name 'Port'.\nEmit a corrected JSON plan for the original request. Same OS and shell."},
        {"role": "assistant", "content": '{"commands":["Get-NetTCPConnection -LocalPort 3000"],"dialect":"powershell"}'}
    ]
    dataset.append({"messages": msg_repair_port})

# 4. Standard PowerShell operations (no tools needed)
powershell_basics = [
    ("list files", ["Get-ChildItem"]),
    ("find text files", ["Get-ChildItem -Filter *.txt"]),
    ("search for error in logs", ["Get-ChildItem -Filter *.log -Recurse | Select-String -Pattern 'error'"]),
    ("show env path", ["$env:PATH"]),
    ("current directory", ["Get-Location"])
]

for _ in range(50):
    for query, cmds in powershell_basics:
        add_example(query, [], cmds, "windows", "powershell")


# Write out the dataset
random.shuffle(dataset)

split_idx = int(len(dataset) * 0.9)
train_data = dataset[:split_idx]
eval_data = dataset[split_idx:]

with open("data/nsh_train.jsonl", "w", encoding="utf-8") as f:
    for d in train_data:
        f.write(json.dumps(d) + "\n")

with open("data/nsh_eval.jsonl", "w", encoding="utf-8") as f:
    for d in eval_data:
        f.write(json.dumps(d) + "\n")

print(f"Generated {len(train_data)} training and {len(eval_data)} evaluation examples.")
