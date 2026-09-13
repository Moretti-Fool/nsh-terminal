from unsloth import FastLanguageModel

model, tokenizer = FastLanguageModel.from_pretrained(
    model_name = "output/nsh-qwen-v2-lora",
    max_seq_length = 512,
    dtype = None,
    load_in_4bit = False,
    device_map = "cpu",
)

model.save_pretrained_gguf("output", tokenizer, quantization_method="q4_k_m")
print("GGUF Export Finished!")
