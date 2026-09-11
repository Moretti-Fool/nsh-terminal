import json
from transformers import AutoTokenizer
t = AutoTokenizer.from_pretrained('Qwen/Qwen2.5-1.5B-Instruct')
msgs = [{'role': 'user', 'content': 'hi'}, {'role': 'assistant', 'tool_calls': [{'type':'function', 'function':{'name': 'run_command', 'arguments': '{"command": "ls"}'}}]}]
print(t.apply_chat_template(msgs, tokenize=False))
