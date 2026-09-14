import os
import torch
from datasets import load_dataset
from trl import SFTTrainer
from transformers import TrainingArguments
from unsloth import FastLanguageModel

# ---------------------------------------------------------
# nsh Fine-Tuning Script (Unsloth)
# Run this in a WSL or Linux environment with a GPU.
# ---------------------------------------------------------

max_seq_length = 2048
dtype = None # Auto detection
load_in_4bit = True # Use 4bit quantization to reduce memory usage

def format_prompt(example):
    # Formats the dataset into ChatML / Qwen format
    messages = example["messages"]
    text = ""
    for msg in messages:
        role = msg["role"]
        content = msg["content"]
        text += f"<|im_start|>{role}\n{content}<|im_end|>\n"
    text += "<|im_start|>assistant\n"
    return {"text": text}

def main():
    print("Loading Base Model (Qwen2.5-Coder-3B)...")
    # We use Qwen2.5-Coder-3B as the base model.
    # Note: You cannot artificially increase a 1.5B model to 3B during fine-tuning.
    # Model architecture (layers, heads, dimensions) is fixed at pre-training.
    model, tokenizer = FastLanguageModel.from_pretrained(
        model_name = "Qwen/Qwen2.5-Coder-3B-Instruct",
        max_seq_length = max_seq_length,
        dtype = dtype,
        load_in_4bit = load_in_4bit,
    )

    print("Applying LoRA Adapters...")
    model = FastLanguageModel.get_peft_model(
        model,
        r = 16, 
        target_modules = ["q_proj", "k_proj", "v_proj", "o_proj",
                          "gate_proj", "up_proj", "down_proj"],
        lora_alpha = 16,
        lora_dropout = 0,
        bias = "none",
        use_gradient_checkpointing = "unsloth",
        random_state = 3407,
        use_rslora = False,
        loftq_config = None,
    )

    print("Loading Dataset...")
    dataset = load_dataset("json", data_files="../training/data/nsh_train_advanced.jsonl", split="train")
    dataset = dataset.map(format_prompt)

    print("Initializing Trainer...")
    trainer = SFTTrainer(
        model = model,
        tokenizer = tokenizer,
        train_dataset = dataset,
        dataset_text_field = "text",
        max_seq_length = max_seq_length,
        dataset_num_proc = 2,
        packing = False,
        args = TrainingArguments(
            per_device_train_batch_size = 2,
            gradient_accumulation_steps = 4,
            warmup_steps = 5,
            num_train_epochs = 3,
            learning_rate = 2e-4,
            fp16 = not torch.cuda.is_bf16_supported(),
            bf16 = torch.cuda.is_bf16_supported(),
            logging_steps = 1,
            optim = "adamw_8bit",
            weight_decay = 0.01,
            lr_scheduler_type = "linear",
            seed = 3407,
            output_dir = "outputs",
        ),
    )

    print("Starting Fine-Tuning...")
    trainer.train()

    print("Exporting Model to GGUF (for Ollama)...")
    # Export to GGUF format for direct Ollama integration
    model.save_pretrained_gguf("nsh_model_3b", tokenizer, quantization_method = "q4_k_m")
    print("Done! Model saved to nsh_model_3b.gguf")

if __name__ == '__main__':
    main()
