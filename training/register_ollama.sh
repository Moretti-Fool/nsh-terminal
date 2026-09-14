#!/bin/bash
cd "$(dirname "$0")"
ollama create nsh-local -f Modelfile
echo "Model registered! Update your nsh config:"
echo '  generation_model = "nsh-local"'
echo '  classifier_model = "nsh-local"'
