from trl import SFTTrainer, SFTConfig
from unsloth import FastLanguageModel
from datasets import Dataset
model, tokenizer = FastLanguageModel.from_pretrained('Qwen/Qwen2.5-1.5B-Instruct', load_in_4bit=True)
ds = Dataset.from_dict({'input_ids': [[1, 2, 3]]})
args = SFTConfig(output_dir='tmp', max_length=512, dataset_kwargs={'skip_prepare_dataset': True}, eos_token='')
trainer = SFTTrainer(model=model, tokenizer=tokenizer, train_dataset=ds, args=args)
print('SUCCESS')