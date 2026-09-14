import requests
import json

OLLAMA_URL = "http://localhost:11434/api/generate"
# We will use the tiny 1.5B model to prove it works on low-end hardware
MODEL = "qwen2.5:1.5b"

# Simulated commands and their outputs
test_cases = [
    {
        "cmd": "git commit -m 'update'",
        "out": "[main 1234abc] update\n 1 file changed, 1 insertion(+)"
    },
    {
        "cmd": "swa deploy",
        "out": "Welcome to Azure Static Web Apps CLI\nDeploying front-end application..."
    },
    {
        "cmd": "psql -U admin",
        "out": "psql (PostgreSQL) 14.2\nType 'help' for help."
    },
    {
        "cmd": "docker run -d nginx",
        "out": "Unable to find image 'nginx:latest' locally\nStatus: Downloaded newer image for nginx:latest"
    },
    {
        "cmd": "git pull origin main",
        "out": "From https://github.com/repo\n * branch main -> FETCH_HEAD\nAlready up to date."
    },
    {
        "cmd": "npm install react",
        "out": "added 1 package, and audited 2 packages in 3s"
    },
    {
        "cmd": "kubectl get pods",
        "out": "NAME                     READY   STATUS    RESTARTS   AGE\nnginx-6799fc88d8-abcde   1/1     Running   0          12m"
    }
]

existing_categories = []
results = []

print(f"--- Stress Testing Auto-Categorization on {MODEL} ---")

for case in test_cases:
    prompt = f"""You categorize terminal commands based on their output.
Existing categories: {json.dumps(existing_categories)}

Command: {case['cmd']}
Command Output: {case['out']}

If the command clearly belongs in one of the existing categories, output the exact existing category name.
If it is a new domain, invent ONE new broad category name (max 2 words, Title Case).
Respond in STRICT JSON format: {{"category": "Name"}}
"""
    
    try:
        r = requests.post(OLLAMA_URL, json={
            "model": MODEL,
            "prompt": prompt,
            "format": "json",
            "stream": False,
            "options": {"temperature": 0.1, "num_predict": 50}
        }, timeout=30)
        
        resp = r.json()
        data = json.loads(resp.get("response", "{}"))
        category = data.get("category", "Unknown").strip()
        
        # Add to known categories if new to simulate dynamic memory
        if category not in existing_categories and category != "Unknown":
            existing_categories.append(category)
            
        results.append({"cmd": case['cmd'], "category": category})
        print(f"[{category.upper()}] <- {case['cmd']}")
        
    except Exception as e:
        print(f"Error on {case['cmd']}: {e}")

print("\n--- Final Discovered Taxonomy ---")
for c in existing_categories:
    print(f"- {c}")
