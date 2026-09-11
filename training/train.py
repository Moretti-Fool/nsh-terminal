import os
import torch
from datasets import load_dataset
from trl import SFTTrainer, SFTConfig
from unsloth import FastLanguageModel

max_length = 512
dtype = None
load_in_4bit = True

print("Loading base model...")
model, tokenizer = FastLanguageModel.from_pretrained(
    model_name = "Qwen/Qwen2.5-1.5B-Instruct",
    dtype = dtype,
    load_in_4bit = load_in_4bit,
)

tokenizer.pad_token = tokenizer.eos_token
tokenizer.pad_token_id = tokenizer.eos_token_id

print("Configuring LoRA...")
model = FastLanguageModel.get_peft_model(
    model,
    r = 32,
    target_modules = ["q_proj", "k_proj", "v_proj", "o_proj",
                      "gate_proj", "up_proj", "down_proj",],
    lora_alpha = 16,
    lora_dropout = 0.05,
    bias = "none",
    use_gradient_checkpointing = "unsloth",
    random_state = 3407,
    use_rslora = False,
    loftq_config = None,
)

print("Loading and preparing datasets...")
train_dataset = load_dataset("json", data_files="data/nsh_train.jsonl", split="train")
eval_dataset = load_dataset("json", data_files="data/nsh_eval.jsonl", split="train")

def format_chat_template(examples):
    formatted_texts = []
    for msgs in examples["messages"]:
        text = tokenizer.apply_chat_template(msgs, tokenize=False, add_generation_prompt=False)
        formatted_texts.append(text)
    return {"input_ids": tokenizer(formatted_texts, truncation=True, max_length=512)["input_ids"]}

train_dataset = train_dataset.map(format_chat_template, batched=True)
eval_dataset = eval_dataset.map(format_chat_template, batched=True)

print("Configuring training...")
args = SFTConfig(
    max_length=512,
    dataset_text_field='text',
    eos_token='<|im_end|>',
    packing=False,
    per_device_train_batch_size = 1,
    gradient_accumulation_steps = 8,
    warmup_ratio = 0.1,
    num_train_epochs = 3,
    learning_rate = 2e-4,
    fp16 = True,
    bf16 = False,
    logging_steps = 10,
    optim = "adamw_8bit",
    weight_decay = 0.01,
    lr_scheduler_type = "linear",
    seed = 3407,
    output_dir = "output/checkpoints",
    save_steps=200,
    eval_strategy="epoch",
)

trainer = SFTTrainer(
    model = model,
    processing_class = tokenizer,
    train_dataset = train_dataset,
    eval_dataset = eval_dataset,
    args = args,
)

print("Starting training...")
trainer_stats = trainer.train()

print("Saving LoRA adapter...")
model.save_pretrained("output/nsh-qwen-lora")
tokenizer.save_pretrained("output/nsh-qwen-lora")

print("Merging LoRA back into base model and saving...")
model.save_pretrained_merged("output/nsh-qwen-merged", tokenizer, save_method = "merged_16bit")
print("Done!")
