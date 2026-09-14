import json
import random
import os

random.seed(42)

# Templates for data generation
powershell_templates = [
    {
        "queries": [
            "list files",
            "show me all files here",
            "what is in this directory",
            "ls",
            "get directory contents"
        ],
        "commands": ["Get-ChildItem"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "list files in {dir}",
            "show contents of {dir}",
            "what's inside {dir}",
            "give me files from {dir}"
        ],
        "commands": ["Get-ChildItem -Path {dir}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "find all text files",
            "list .txt files",
            "show text files",
            "give me all txt files in this folder"
        ],
        "commands": ["Get-ChildItem -Filter *.txt"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "show me all running processes",
            "list active system processes",
            "running procs",
            "I need to see what processes are running",
            "processes"
        ],
        "commands": ["Get-Process"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "kill process with id {pid}",
            "stop proc {pid}",
            "terminate process {pid}",
            "end task {pid}"
        ],
        "commands": ["Stop-Process -Id {pid} -Force"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "kill process named {pname}",
            "stop the {pname} process",
            "terminate {pname}"
        ],
        "commands": ["Stop-Process -Name {pname} -Force"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "show system info",
            "get computer info",
            "what are my system specs"
        ],
        "commands": ["Get-ComputerInfo"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "show active network connections",
            "netstat",
            "list tcp connections"
        ],
        "commands": ["Get-NetTCPConnection"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "ping google",
            "test connection to google.com",
            "is the internet working"
        ],
        "commands": ["Test-NetConnection -ComputerName google.com"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "download {url} to {file}",
            "fetch {url} and save as {file}",
            "wget {url}"
        ],
        "commands": ["Invoke-WebRequest -Uri {url} -OutFile {file}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "check git status",
            "what changed in git",
            "git status"
        ],
        "commands": ["git status"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "commit with message {msg}",
            "save changes as {msg}",
            "git commit {msg}"
        ],
        "commands": ["git commit -m \"{msg}\""],
        "dialect": "powershell"
    },
    {
        "queries": [
            "find large files",
            "find all files over 100MB",
            "files bigger than 100mb"
        ],
        "commands": ["Get-ChildItem -Recurse | Where-Object { $_.Length -gt 100MB }"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "search for {text} in files",
            "grep {text}",
            "find string {text}"
        ],
        "commands": ["Select-String -Pattern \"{text}\" -Path *"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "zip the folder {dir}",
            "compress {dir} to archive",
            "create zip from {dir}"
        ],
        "commands": ["Compress-Archive -Path {dir} -DestinationPath archive.zip"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "show environment path",
            "what is my path variable",
            "echo path"
        ],
        "commands": ["$env:PATH"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "create a new directory {dir}",
            "make dir {dir}",
            "mkdir {dir}"
        ],
        "commands": ["New-Item -ItemType Directory -Path {dir}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "delete file {file}",
            "remove {file}",
            "del {file}",
            "rm {file}"
        ],
        "commands": ["Remove-Item -Path {file}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "read file {file}",
            "cat {file}",
            "show contents of {file}"
        ],
        "commands": ["Get-Content -Path {file}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "list docker containers",
            "docker ps",
            "what containers are running"
        ],
        "commands": ["docker ps"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "show docker logs for {container}",
            "logs for {container}"
        ],
        "commands": ["docker logs {container}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "check file permissions for {file}",
            "who owns {file}",
            "get acl {file}"
        ],
        "commands": ["Get-Acl -Path {file}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "build docker image {img}",
            "docker build {img}"
        ],
        "commands": ["docker build -t {img} ."],
        "dialect": "powershell"
    },
    {
        "queries": [
            "sort files by size",
            "list files ordered by size",
            "biggest files first"
        ],
        "commands": ["Get-ChildItem | Sort-Object Length -Descending"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "copy {file} to {dir}",
            "cp {file} {dir}"
        ],
        "commands": ["Copy-Item -Path {file} -Destination {dir}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "move {file} to {dir}",
            "mv {file} {dir}"
        ],
        "commands": ["Move-Item -Path {file} -Destination {dir}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "write {text} to {file}",
            "echo {text} > {file}"
        ],
        "commands": ["Set-Content -Path {file} -Value \"{text}\""],
        "dialect": "powershell"
    },
    {
        "queries": [
            "check if {file} exists",
            "does {file} exist"
        ],
        "commands": ["Test-Path -Path {file}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "list services",
            "show all services",
            "services"
        ],
        "commands": ["Get-Service"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "start service {svc}",
            "run service {svc}"
        ],
        "commands": ["Start-Service -Name {svc}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "stop service {svc}",
            "kill service {svc}"
        ],
        "commands": ["Stop-Service -Name {svc}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "check event log",
            "show errors in event log"
        ],
        "commands": ["Get-EventLog -LogName System -EntryType Error"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "resolve dns for {domain}",
            "nslookup {domain}",
            "get ip of {domain}"
        ],
        "commands": ["Resolve-DnsName -Name {domain}"],
        "dialect": "powershell"
    },
    {
        "queries": [
            "unzip {file}",
            "extract archive {file}"
        ],
        "commands": ["Expand-Archive -Path {file} -DestinationPath ."],
        "dialect": "powershell"
    },
    {
        "queries": [
            "count lines in {file}",
            "how many lines in {file}"
        ],
        "commands": ["Get-Content {file} | Measure-Object -Line"],
        "dialect": "powershell"
    }
]

bash_templates = [
    {
        "queries": ["list files", "ls", "what is in this dir", "show directory"],
        "commands": ["ls -la"],
        "dialect": "bash"
    },
    {
        "queries": ["copy {file} to {dir}", "cp {file} to {dir}"],
        "commands": ["cp {file} {dir}/"],
        "dialect": "bash"
    },
    {
        "queries": ["move {file} to {dir}", "mv {file} to {dir}"],
        "commands": ["mv {file} {dir}/"],
        "dialect": "bash"
    },
    {
        "queries": ["delete {file}", "rm {file}", "remove {file}"],
        "commands": ["rm {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["create directory {dir}", "mkdir {dir}", "make dir {dir}"],
        "commands": ["mkdir -p {dir}"],
        "dialect": "bash"
    },
    {
        "queries": ["read {file}", "cat {file}", "show {file} contents"],
        "commands": ["cat {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["first 10 lines of {file}", "head {file}", "start of {file}"],
        "commands": ["head -n 10 {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["last 10 lines of {file}", "tail {file}", "end of {file}"],
        "commands": ["tail -n 10 {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["find files named {text}", "search for {text} file"],
        "commands": ["find . -name \"*{text}*\""],
        "dialect": "bash"
    },
    {
        "queries": ["change permissions of {file} to {perms}", "chmod {perms} {file}"],
        "commands": ["chmod {perms} {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["change owner of {file} to {user}", "chown {user} {file}"],
        "commands": ["chown {user} {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["show running processes", "ps", "top", "list processes"],
        "commands": ["ps aux"],
        "dialect": "bash"
    },
    {
        "queries": ["kill process {pid}", "stop {pid}", "end {pid}"],
        "commands": ["kill -9 {pid}"],
        "dialect": "bash"
    },
    {
        "queries": ["check system status", "systemctl status", "system health"],
        "commands": ["systemctl status"],
        "dialect": "bash"
    },
    {
        "queries": ["view system logs", "journalctl", "syslog"],
        "commands": ["journalctl -xe"],
        "dialect": "bash"
    },
    {
        "queries": ["download {url}", "curl {url}", "wget {url}"],
        "commands": ["curl -O {url}"],
        "dialect": "bash"
    },
    {
        "queries": ["ping {domain}", "check connection to {domain}"],
        "commands": ["ping -c 4 {domain}"],
        "dialect": "bash"
    },
    {
        "queries": ["trace route to {domain}", "traceroute {domain}"],
        "commands": ["traceroute {domain}"],
        "dialect": "bash"
    },
    {
        "queries": ["show network stats", "netstat", "ss"],
        "commands": ["ss -tulpn"],
        "dialect": "bash"
    },
    {
        "queries": ["resolve {domain}", "dig {domain}", "nslookup {domain}"],
        "commands": ["dig {domain}"],
        "dialect": "bash"
    },
    {
        "queries": ["find {text} in {file}", "grep {text} {file}", "search {text} inside {file}"],
        "commands": ["grep \"{text}\" {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["replace {text} with {newtext} in {file}", "sed replace {text}"],
        "commands": ["sed -i 's/{text}/{newtext}/g' {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["sort lines in {file}", "sort {file}"],
        "commands": ["sort {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["count lines in {file}", "wc {file}"],
        "commands": ["wc -l {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["extract {file}", "tar extract {file}", "unzip {file}"],
        "commands": ["tar -xzf {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["compress {dir}", "tar create {dir}"],
        "commands": ["tar -czf archive.tar.gz {dir}"],
        "dialect": "bash"
    },
    {
        "queries": ["install {pkg}", "apt install {pkg}"],
        "commands": ["sudo apt-get install {pkg}"],
        "dialect": "bash"
    },
    {
        "queries": ["update packages", "apt update"],
        "commands": ["sudo apt-get update"],
        "dialect": "bash"
    },
    {
        "queries": ["git status", "check git"],
        "commands": ["git status"],
        "dialect": "bash"
    },
    {
        "queries": ["git pull", "update from remote"],
        "commands": ["git pull"],
        "dialect": "bash"
    },
    {
        "queries": ["git push", "upload changes"],
        "commands": ["git push"],
        "dialect": "bash"
    },
    {
        "queries": ["git log", "show commit history"],
        "commands": ["git log --oneline"],
        "dialect": "bash"
    },
    {
        "queries": ["docker ps", "list docker containers"],
        "commands": ["docker ps"],
        "dialect": "bash"
    },
    {
        "queries": ["docker logs {container}", "logs for {container}"],
        "commands": ["docker logs {container}"],
        "dialect": "bash"
    },
    {
        "queries": ["find files modified in last 7 days"],
        "commands": ["find . -mtime -7"],
        "dialect": "bash"
    },
    {
        "queries": ["awk print second column of {file}"],
        "commands": ["awk '{{print $2}}' {file}"],
        "dialect": "bash"
    },
    {
        "queries": ["run {cmd} on all files"],
        "commands": ["find . -type f | xargs {cmd}"],
        "dialect": "bash"
    }
]

cmd_templates = [
    {
        "queries": ["list files", "dir", "show directory"],
        "commands": ["dir"],
        "dialect": "cmd"
    },
    {
        "queries": ["copy {file} to {dir}", "cp {file} to {dir}"],
        "commands": ["copy {file} {dir}"],
        "dialect": "cmd"
    },
    {
        "queries": ["move {file} to {dir}", "mv {file} {dir}"],
        "commands": ["move {file} {dir}"],
        "dialect": "cmd"
    },
    {
        "queries": ["delete {file}", "del {file}"],
        "commands": ["del {file}"],
        "dialect": "cmd"
    },
    {
        "queries": ["make directory {dir}", "mkdir {dir}"],
        "commands": ["mkdir {dir}"],
        "dialect": "cmd"
    },
    {
        "queries": ["read {file}", "type {file}", "show {file} contents"],
        "commands": ["type {file}"],
        "dialect": "cmd"
    },
    {
        "queries": ["search for {text} in {file}", "findstr {text} {file}"],
        "commands": ["findstr \"{text}\" {file}"],
        "dialect": "cmd"
    },
    {
        "queries": ["sort {file}"],
        "commands": ["sort {file}"],
        "dialect": "cmd"
    },
    {
        "queries": ["list network shares", "net use"],
        "commands": ["net use"],
        "dialect": "cmd"
    },
    {
        "queries": ["check service {svc}", "sc query {svc}"],
        "commands": ["sc query {svc}"],
        "dialect": "cmd"
    },
    {
        "queries": ["stop service {svc}", "net stop {svc}"],
        "commands": ["net stop {svc}"],
        "dialect": "cmd"
    },
    {
        "queries": ["start service {svc}", "net start {svc}"],
        "commands": ["net start {svc}"],
        "dialect": "cmd"
    },
    {
        "queries": ["loop through files and echo names", "for files echo"],
        "commands": ["for %f in (*) do echo %f"],
        "dialect": "cmd"
    },
    {
        "queries": ["find files modified today", "forfiles today"],
        "commands": ["forfiles /D +0"],
        "dialect": "cmd"
    },
    {
        "queries": ["clear screen", "cls"],
        "commands": ["cls"],
        "dialect": "cmd"
    },
    {
        "queries": ["show path", "echo path"],
        "commands": ["echo %PATH%"],
        "dialect": "cmd"
    },
    {
        "queries": ["ipconfig", "show ip"],
        "commands": ["ipconfig /all"],
        "dialect": "cmd"
    },
    {
        "queries": ["check disk", "chkdsk"],
        "commands": ["chkdsk C:"],
        "dialect": "cmd"
    },
    {
        "queries": ["tree view", "show folder tree"],
        "commands": ["tree /F"],
        "dialect": "cmd"
    },
    {
        "queries": ["who am i", "current user"],
        "commands": ["whoami"],
        "dialect": "cmd"
    }
]

# Random value lists to populate templates
values = {
    "dir": ["documents", "downloads", "src", "app/data", "build", "empty_dir", "folder with spaces"],
    "file": ["config.json", "data.csv", "index.html", "main.py", "script.ps1", "file with space.txt", "very_long_file_name.txt"],
    "text": ["error", "warning", "exception", "TODO", "password", "127.0.0.1", "function main()"],
    "newtext": ["debug", "info", "FIXME", "secret"],
    "pid": ["1234", "5678", "9012", "4321"],
    "pname": ["chrome", "svchost", "node", "python", "explorer"],
    "url": ["https://api.github.com", "http://localhost:8080", "https://google.com"],
    "domain": ["google.com", "localhost", "github.com", "1.1.1.1"],
    "msg": ["Initial commit", "Fix bug", "Update docs", "WIP"],
    "container": ["web", "db", "redis", "nginx_1"],
    "img": ["ubuntu:latest", "node:18", "alpine"],
    "perms": ["755", "644", "777", "+x"],
    "user": ["root", "admin", "guest", "www-data"],
    "pkg": ["curl", "git", "vim", "htop"],
    "svc": ["spooler", "wuauserv", "docker", "ssh"]
}

def format_string(text, kwargs):
    for k, v in kwargs.items():
        text = text.replace("{" + k + "}", str(v))
    return text

def generate_multi(templates, dialect):
    # pick 2 or 3 random templates
    num = random.randint(2, 3)
    chosen = random.sample(templates, num)
    
    q_parts = []
    c_parts = []
    
    for t in chosen:
        q = random.choice(t["queries"])
        c = t["commands"][0]
        
        # Format them
        import re
        keys = set(re.findall(r'\{([a-z]+)\}', q + c))
        kwargs = {k: random.choice(values.get(k, ["val"])) for k in keys}
        
        q_parts.append(format_string(q, kwargs))
        c_parts.append(format_string(c, kwargs))
    
    join_words = [" and ", ", then ", ". After that, "]
    query = random.choice(join_words).join(q_parts)
    
    return {
        "queries": [query],
        "commands": c_parts,
        "dialect": dialect
    }

def generate_item(templates, dialect, is_multi=False):
    if is_multi:
        t = generate_multi(templates, dialect)
        q = t["queries"][0]
        c = t["commands"]
    else:
        t = random.choice(templates)
        q = random.choice(t["queries"])
        c = list(t["commands"])
        
        # Format
        import re
        keys = set(re.findall(r'\{([a-z]+)\}', q + "".join(c)))
        kwargs = {k: random.choice(values.get(k, ["val"])) for k in keys}
        
        q = format_string(q, kwargs)
        c = [format_string(cmd, kwargs) for cmd in c]
        
    # Generate system prompt
    oss = ["windows", "linux", "darwin"]
    cwd = random.choice(["/home/user", "C:\\Users\\Admin", "/var/www/html", "D:\\Projects"])
    
    sys_prompt = f"You are a helpful assistant that translates natural language into shell commands. The current operating system is {random.choice(oss)}, the shell is {dialect}, and the current working directory is {cwd}. Return a JSON object with 'commands' (array of strings) and 'dialect' (string representing the shell dialect). Do not provide any other output."
    
    return {
        "messages": [
            {"role": "system", "content": sys_prompt},
            {"role": "user", "content": q},
            {"role": "assistant", "content": json.dumps({"commands": c, "dialect": dialect})}
        ]
    }

def main():
    target_total = 5000
    
    pwsh_count = int(target_total * 0.50)
    bash_count = int(target_total * 0.30)
    cmd_count = int(target_total * 0.20)
    
    dataset = []
    
    for _ in range(pwsh_count):
        is_multi = random.random() < 0.1
        dataset.append(generate_item(powershell_templates, "powershell", is_multi))
        
    for _ in range(bash_count):
        is_multi = random.random() < 0.1
        dataset.append(generate_item(bash_templates, "bash", is_multi))
        
    for _ in range(cmd_count):
        is_multi = random.random() < 0.1
        dataset.append(generate_item(cmd_templates, "cmd", is_multi))
        
    # Shuffle dataset
    random.shuffle(dataset)
    
    # Split 90/10
    split_idx = int(len(dataset) * 0.9)
    train_data = dataset[:split_idx]
    eval_data = dataset[split_idx:]
    
    # Ensure dirs
    os.makedirs("C:/Users/sanch/OneDrive/Documents/nsh/training/data", exist_ok=True)
    
    # Validate and write
    def write_jsonl(path, data):
        with open(path, 'w', encoding='utf-8') as f:
            for item in data:
                # Validate JSON structure
                assert "messages" in item
                assert len(item["messages"]) == 3
                
                # Check assistant content is valid JSON string
                try:
                    json.loads(item["messages"][2]["content"])
                except json.JSONDecodeError:
                    raise ValueError(f"Invalid JSON in assistant response: {item['messages'][2]['content']}")
                    
                f.write(json.dumps(item) + "\n")
                
    write_jsonl("C:/Users/sanch/OneDrive/Documents/nsh/training/data/nsh_train.jsonl", train_data)
    write_jsonl("C:/Users/sanch/OneDrive/Documents/nsh/training/data/nsh_eval.jsonl", eval_data)
    
    # Create gitkeep
    with open("C:/Users/sanch/OneDrive/Documents/nsh/training/data/.gitkeep", 'w') as f:
        pass
        
    print(f"Generated {len(dataset)} total examples.")
    print(f"Train split: {len(train_data)} examples.")
    print(f"Eval split: {len(eval_data)} examples.")
    
if __name__ == "__main__":
    main()
