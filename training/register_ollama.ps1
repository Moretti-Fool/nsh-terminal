Set-Location $PSScriptRoot
ollama create nsh-local -f Modelfile
Write-Host "Model registered! Update your nsh config:"
Write-Host '  generation_model = "nsh-local"'
Write-Host '  classifier_model = "nsh-local"'
