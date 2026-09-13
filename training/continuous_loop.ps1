Write-Host "========================================="
Write-Host "   NSH CONTINUOUS TRAINING LOOP START    "
Write-Host "========================================="

$max_loops = 5

for ($i = 1; $i -le $max_loops; $i++) {
    Write-Host "`n---[ LOOP $i ]---"
    
    # 1. Augment the dataset with fresh diverse examples
    Write-Host "Augmenting dataset using 7B model..."
    .\venv\Scripts\python.exe augment_dataset.py

    # 2. Run the fine-tuning loop
    Write-Host "Starting fine-tuning step..."
    .\venv\Scripts\python.exe train.py

    # 3. Export the model
    Write-Host "Exporting to GGUF..."
    .\venv\Scripts\python.exe export_gguf.py

    # 4. Reload into Ollama
    Write-Host "Updating local Ollama model..."
    ollama create nsh-local:v2 -f Modelfile

    # 5. Modify train.py for the NEXT loop to train on its own merged checkpoint!
    Write-Host "Updating train.py to use the merged checkpoint for Loop $($i + 1)..."
    (Get-Content train.py) -replace 'model_name = ".*"', 'model_name = "output/nsh-qwen-v2-merged"' | Set-Content train.py

    Write-Host "Loop $i complete!"
}

Write-Host "========================================="
Write-Host "      CONTINUOUS TRAINING FINISHED       "
Write-Host "========================================="
