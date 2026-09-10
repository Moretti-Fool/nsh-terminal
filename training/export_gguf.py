import os
import subprocess
import sys
import shutil

def main():
    base_dir = os.path.dirname(os.path.abspath(__file__))
    output_dir = os.path.join(base_dir, 'output')
    llama_cpp_dir = os.path.join(base_dir, 'llama.cpp')
    merged_model_dir = os.path.join(output_dir, 'nsh-qwen-merged')
    
    if not os.path.exists(llama_cpp_dir):
        print("Cloning llama.cpp repository...")
        subprocess.run(['git', 'clone', 'https://github.com/ggerganov/llama.cpp', llama_cpp_dir], check=True)

    f16_gguf_path = os.path.join(output_dir, 'nsh-qwen-f16.gguf')
    q4km_gguf_path = os.path.join(output_dir, 'nsh-qwen-q4km.gguf')

    # Convert to GGUF f16
    convert_script = os.path.join(llama_cpp_dir, 'convert_hf_to_gguf.py')
    print("Converting to GGUF f16...")
    subprocess.run([sys.executable, convert_script, merged_model_dir, '--outfile', f16_gguf_path, '--outtype', 'f16'], check=True)

    # Quantize to Q4_K_M
    quantize_bin = os.path.join(llama_cpp_dir, 'build', 'bin', 'llama-quantize')
    
    # On Windows it might be an .exe
    if os.name == 'nt' and not os.path.exists(quantize_bin):
        quantize_bin += '.exe'

    print(f"Quantizing to Q4_K_M using {quantize_bin}...")
    subprocess.run([quantize_bin, f16_gguf_path, q4km_gguf_path, 'Q4_K_M'], check=True)

    if os.path.exists(q4km_gguf_path):
        size_bytes = os.path.getsize(q4km_gguf_path)
        size_mb = size_bytes / (1024 * 1024)
        print(f"\nQuantization complete!")
        print(f"Output file: {q4km_gguf_path}")
        print(f"File size: {size_mb:.2f} MB")
        print("\nNext steps:")
        print("1. Run the Ollama registration script (register_ollama.sh or register_ollama.ps1)")
        print("2. Test the model using evaluate.py")
    else:
        print("Error: Output file not generated.")

if __name__ == "__main__":
    main()
