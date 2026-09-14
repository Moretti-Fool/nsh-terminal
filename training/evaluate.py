import json
import requests
import sys

def main():
    eval_file = 'data/nsh_eval.jsonl'
    url = 'http://localhost:11434/api/generate'
    model = 'nsh-local'
    
    total = 0
    parseable = 0
    schema_ok = 0
    exact_match = 0
    dialect_match = 0
    
    try:
        with open(eval_file, 'r', encoding='utf-8') as f:
            for line in f:
                if not line.strip():
                    continue
                total += 1
                data = json.loads(line)
                
                # Infer format (either direct input/expected or messages array)
                if 'messages' in data:
                    user_msg = next((m['content'] for m in data['messages'] if m['role'] == 'user'), '')
                    ast_msg = next((m['content'] for m in data['messages'] if m['role'] == 'assistant'), '{}')
                    expected = json.loads(ast_msg)
                else:
                    user_msg = data.get('input', '')
                    expected = data.get('expected', {})

                payload = {
                    "model": model,
                    "prompt": user_msg,
                    "stream": False
                }
                
                resp = requests.post(url, json=payload)
                if resp.status_code != 200:
                    print(f"Error calling ollama: {resp.text}")
                    continue
                    
                resp_data = resp.json()
                gen_text = resp_data.get('response', '').strip()
                
                try:
                    gen_json = json.loads(gen_text)
                    parseable += 1
                    
                    if isinstance(gen_json, dict) and 'commands' in gen_json and 'dialect' in gen_json:
                        schema_ok += 1
                        
                        expected_commands = expected.get('commands', [])
                        expected_dialect = expected.get('dialect', '')
                        
                        if gen_json.get('dialect') == expected_dialect:
                            dialect_match += 1
                            
                        if gen_json.get('commands') == expected_commands:
                            exact_match += 1
                            
                except json.JSONDecodeError:
                    pass

        if total > 0:
            print(f"Total evaluated: {total}")
            print(f"Parseable JSON: {parseable/total*100:.2f}%")
            print(f"Schema Adherence: {schema_ok/total*100:.2f}%")
            print(f"Dialect Accuracy: {dialect_match/total*100:.2f}%")
            print(f"Exact Match: {exact_match/total*100:.2f}%")
        else:
            print("No examples found to evaluate.")

    except FileNotFoundError:
        print(f"Eval file not found: {eval_file}")

if __name__ == "__main__":
    main()
