import requests
import os
import time

class OpenRouterAdapter:
    """
    Central adapter for connecting the TLE to the Genius Cloud (OpenRouter).
    Handles rate limiting, error recovery, and automatic model failover.
    """
    def __init__(self, env_path=".env"):
        self.api_key = self._load_key(env_path)
        self.base_url = "https://openrouter.ai/api/v1/chat/completions"
        
        # MODEL POOL: Prioritized list of high-quality free models
        self.model_pool = [
          
             "openai/gpt-oss-120b:free",                 # strong fallback
    "openrouter/owl-alpha",                     # worked in our last run
    "meta-llama/llama-3.3-70b-instruct:free",  # strong when available
        "openrouter/free",         
        "meta-llama/llama-3.3-70b-instruct:free",
    "openai/gpt-oss-120b:free",
    "arcee-ai/trinity-large-thinking:free",
    "deepseek/deepseek-v4-flash:free",
    "openrouter/owl-alpha",
    "poolside/laguna-m.1:free",
    "minimax/minimax-m2.5:free",
    "qwen/qwen3-next-80b-a3b-instruct:free",
    "baidu/cobuddy:free",
    "poolside/laguna-m.1:free",
    "google/gemma-4-26b-a4b-it:free",
    "google/gemma-4-31b-it:free",
      "google/gemini-2.0-flash-exp:free",
            "google/gemini-flash-1.5-free",
            "meta-llama/llama-3.1-70b-instruct:free",
            "meta-llama/llama-3.2-3b-instruct:free",
            "mistralai/pixtral-12b:free",
            "qwen/qwen-2.5-72b-instruct:free",
            "microsoft/phi-3-medium-128k-instruct:free",
        ]
        self.current_model_idx = 0

    def _load_key(self, path):
        if not os.path.exists(path):
            return None
        with open(path, "r", encoding="utf-8") as f:
            for line in f:
                if line.startswith("OPENROUTER_API_KEY="):
                    return line.split("=")[1].strip()
        return None

    def ask(self, prompt, system_prompt="You are a Master AI Architect.", model=None):
        if not self.api_key:
            return "Error: OPENROUTER_API_KEY not found in .env"

        # If a specific model is requested, use it; otherwise use the pool
        models_to_try = [model] if model else self.model_pool[self.current_model_idx:]

        headers = {
            "Authorization": f"Bearer {self.api_key}",
            "Content-Type": "application/json"
        }

        for target_model in models_to_try:
            payload = {
                "model": target_model,
                "messages": [
                    {"role": "system", "content": system_prompt},
                    {"role": "user", "content": prompt}
                ],
                "temperature": 0.1
            }

            # Retry logic for the specific model
            for attempt in range(2):
                try:
                    resp = requests.post(self.base_url, headers=headers, json=payload, timeout=60)
                    if resp.status_code == 200:
                        return resp.json()['choices'][0]['message']['content'].strip()
                    
                    elif resp.status_code == 429:
                        print(f"[OpenRouter] Model {target_model} rate limited. (Attempt {attempt+1})")
                        if attempt == 0:
                            print("Slowing down... sleeping 5s.")
                            time.sleep(5)
                            continue
                        else:
                            print("Switching to next model in pool...")
                            self.current_model_idx = (self.current_model_idx + 1) % len(self.model_pool)
                            break # Move to next model in outer loop

                    else:
                        print(f"[OpenRouter] Model {target_model} error {resp.status_code}. Skipping...")
                        break # Move to next model

                except Exception as e:
                    print(f"[OpenRouter] Connection error with {target_model}: {e}")
                    break

        return "Error: All models in pool exhausted or rate limited."

if __name__ == "__main__":
    adapter = OpenRouterAdapter()
    if adapter.api_key:
        print(f"Key loaded. Model Pool Size: {len(adapter.model_pool)}")
        print(f"Primary Model: {adapter.model_pool[0]}")
        print("Testing Failover Connection...")
        result = adapter.ask("Perform a logic check: is 1+1=2?")
        print(f"Cloud Response: {result}")
    else:
        print("No .env found. Please create .env with OPENROUTER_API_KEY.")
