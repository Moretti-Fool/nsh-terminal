Write-Host "Waiting for generate_dataset.py to finish generating the 315 multi-shell examples..." -ForegroundColor Cyan
while (Get-WmiObject Win32_Process | Where-Object { \.CommandLine -match 'generate_dataset.py' }) {
    Start-Sleep -Seconds 5
}
Write-Host "Dataset generation complete!" -ForegroundColor Green
Write-Host "Starting Unsloth 3B Fine-Tuning..." -ForegroundColor Cyan

cd ../training
& .\venv\Scripts\python.exe train.py
Write-Host "Training finished! Press any key to exit..." -ForegroundColor Green
[Console]::ReadKey()
