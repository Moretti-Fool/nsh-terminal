# Personal Setup Guide

Since you have trained your own highly specialized **nsh-local** model and hosted it securely on your private GitHub repository, here are your personal quick-reference instructions for pulling the model to a new machine and getting your shell up and running.

## 1. Prerequisites
Ensure you have the following installed on your new machine:
- **Go** (for compiling \
sh\)
- **Ollama** (for serving the model locally)
- **GitHub CLI (\gh\)** (for authenticating and pulling the private release)

## 2. Clone the Repository
`powershell
gh auth login
gh repo clone Moretti-Fool/nsh-terminal
cd nsh-terminal
`

## 3. Pull Your Private Model
Because the model is hosted via GitHub Releases on a private repository, you must be authenticated via \gh\:
`powershell
# Create the output directory where the Modelfile expects the model
mkdir training\output

# Download the 940MB GGUF from your private release
gh release download model-v1.0 -p "*.gguf" --dir training\output
`

## 4. Register the Model with Ollama
Your repository includes a \Modelfile\ pre-configured with the exact system prompts, JSON formatting parameters, and temperature settings required for the terminal dialect.

`powershell
# Make sure Ollama is running in the background, then run:
cd training
ollama create nsh-local -f Modelfile
cd ..
`

## 5. Compile and Configure
Build the \
sh\ binary:
`powershell
go build -o nsh.exe main.go
`

The application defaults to \
sh-local\ out of the box now. If you have an old config file lingering in your AppData, make sure to update it:

**Windows**: \%APPDATA%\nsh\config.toml\
**macOS/Linux**: \~/.config/nsh/config.toml\

`	oml
[ollama]
classifier_model = "nsh-local"
generation_model = "nsh-local"
`

## 6. Run it!
Drop the compiled binary into a folder that is registered in your system \PATH\ and type \
sh\ anywhere!
