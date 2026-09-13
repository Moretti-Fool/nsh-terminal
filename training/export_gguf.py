import os
import subprocess
import sys

def main():
    base_dir = os.path.dirname(os.path.abspath(__file__))
    output_dir = os.path.join(base_dir, 'output')
    llama_cpp_dir = os.path.join(base_dir, 'llama.cpp')
    merged_model_dir = os.path.join(output_dir, 'nsh-qwen-v2-merged')
    
    if not os.path.exists(llama_cpp_dir):
        print('Cloning llama.cpp...')
        subprocess.run(['git', 'clone', 'https://github.com/ggerganov/llama.cpp', llama_cpp_dir], check=True)

    f16_gguf_path = os.path.join(output_dir, 'nsh-qwen-v2-f16.gguf')
    q4km_gguf_path = os.path.join(output_dir, 'nsh-qwen-v2-q4km.gguf')

    convert_script = os.path.join(llama_cpp_dir, 'convert_hf_to_gguf.py')
    print("Converting to GGUF f16...")
    subprocess.run([sys.executable, convert_script, merged_model_dir, '--outfile', f16_gguf_path, '--outtype', 'f16'], check=True)

    try:
        print('Building llama-quantize with CMake...')
        subprocess.run(['cmake', '-B', 'build'], cwd=llama_cpp_dir, check=True)
        subprocess.run(['cmake', '--build', 'build', '--config', 'Release', '-j', '--target', 'llama-quantize'], cwd=llama_cpp_dir, check=True)
        
        quantize_bin = os.path.join(llama_cpp_dir, 'build', 'bin', 'llama-quantize')
        if not os.path.exists(quantize_bin):
            quantize_bin = os.path.join(llama_cpp_dir, 'build', 'llama-quantize')
        
        print(f"Quantizing to Q4_K_M using {quantize_bin}...")
        subprocess.run([quantize_bin, f16_gguf_path, q4km_gguf_path, 'Q4_K_M'], check=True)
    except Exception as e:
        print(f"Warning: CMake or quantization failed: {e}")
        print("Falling back to just f16 model.")
        q4km_gguf_path = f16_gguf_path

    if os.path.exists(q4km_gguf_path):
        size_bytes = os.path.getsize(q4km_gguf_path)
        size_mb = size_bytes / (1024 * 1024)
        print(f"\nQuantization complete!")
        print(f"Output file: {q4km_gguf_path}")
        print(f"File size: {size_mb:.2f} MB")
    else:
        print("Error: Output file not generated.")

if __name__ == "__main__":
    main()
