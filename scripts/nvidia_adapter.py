import os
import requests


class NvidiaNimAdapter:
    """
    Elite-level adapter for NVIDIA NIM using OpenAI-compatible client.
    Supports reasoning-enabled Nemotron models.
    """

    def __init__(self, env_path=".env"):
        self.api_key = self._load_key(env_path)
        # self.invoke_url = "https://integrate.api.nvidia.com/v1/"
        self.default_model = "minimaxai/minimax-m2.7"

        # if self.api_key:
        #     self.client = OpenAI(
        #         base_url=self.invoke_url,
        #         api_key=self.api_key
        #     )
        # else:
        #     self.client = None

    def _load_key(self, path):
        if not os.path.exists(path):
            return None
        with open(path, "r") as f:
            for line in f:
                if line.startswith("NVIDIA_API_KEY="):
                    return line.split("=")[1].strip()
        return None

    def ask(self, prompt, system_prompt="You are a Master AI Logic Engineer.", model=None):
        if not self.api_key:
            return "Error: NVIDIA_API_KEY not found in .env"

        model = model or self.default_model

        url = "https://integrate.api.nvidia.com/v1/chat/completions"

        headers = {
            "Authorization": f"Bearer {self.api_key}",
            "Content-Type": "application/json"
        }

        payload = {
        "model": model,
        "messages": [
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": prompt}
        ],
        "temperature": 0.6,
        "top_p": 0.95,
        "max_tokens": 2048,
        "stream": False
        }

        try:
            response = requests.post(url, headers=headers, json=payload, timeout=60)

            print("STATUS:", response.status_code)
            print("RAW:", response.text)

            if response.status_code == 200:
                return response.json()["choices"][0]["message"]["content"]

            return f"FAILED [{response.status_code}]: {response.text}"

        except Exception as e:
            return f"Error connecting to NVIDIA: {e}"

if __name__ == "__main__":
    adapter = NvidiaNimAdapter()

    if adapter.api_key:
        print("NVIDIA Key loaded. Testing reasoning connection...")

        result = adapter.ask(
            "Is the law of non-contradiction a physical or a logical law? Explain in one sentence."
        )

        print(f"NVIDIA Result:\n{result}")

    else:
        print("No .env found or NVIDIA_API_KEY missing.")