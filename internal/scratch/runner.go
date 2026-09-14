package scratch

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"
)

var pythonSignals = []string{
	"pandas", "numpy", "matplotlib", "pyplot", "seaborn",
	"dataframe", "scipy", "sklearn", "scikit",
	"openpyxl", "xlrd", "plotly", "bokeh",
	"with python", "using python", "in python",
	"python script", "python code",
	"using pandas", "with pandas",
	"using numpy", "with numpy",
	"using matplotlib", "with matplotlib",
}

var importToPip = map[string]string{
	"pandas":       "pandas",
	"numpy":        "numpy",
	"matplotlib":   "matplotlib",
	"seaborn":      "seaborn",
	"sklearn":      "scikit-learn",
	"scipy":        "scipy",
	"requests":     "requests",
	"flask":        "flask",
	"fastapi":      "fastapi",
	"uvicorn":      "uvicorn",
	"sqlalchemy":   "sqlalchemy",
	"PIL":          "Pillow",
	"cv2":          "opencv-python",
	"torch":        "torch",
	"tensorflow":   "tensorflow",
	"bs4":          "beautifulsoup4",
	"lxml":         "lxml",
	"yaml":         "pyyaml",
	"dotenv":       "python-dotenv",
	"openpyxl":     "openpyxl",
	"xlrd":         "xlrd",
	"xlsxwriter":   "xlsxwriter",
	"psycopg2":     "psycopg2-binary",
	"pymongo":      "pymongo",
	"redis":        "redis",
	"boto3":        "boto3",
	"paramiko":     "paramiko",
	"cryptography": "cryptography",
	"jwt":          "PyJWT",
	"rich":         "rich",
	"click":        "click",
	"httpx":        "httpx",
	"aiohttp":      "aiohttp",
	"pydantic":     "pydantic",
	"tabulate":     "tabulate",
	"tqdm":         "tqdm",
	"plotly":       "plotly",
	"dash":         "dash",
	"streamlit":    "streamlit",
}

var stdlibModules = map[string]bool{
	"os": true, "sys": true, "json": true, "csv": true,
	"math": true, "random": true, "datetime": true, "time": true,
	"re": true, "io": true, "pathlib": true, "glob": true,
	"shutil": true, "subprocess": true, "threading": true,
	"collections": true, "itertools": true, "functools": true,
	"typing": true, "abc": true, "copy": true, "pprint": true,
	"argparse": true, "logging": true, "unittest": true,
	"hashlib": true, "hmac": true, "base64": true,
	"urllib": true, "http": true, "socket": true,
	"sqlite3": true, "xml": true, "html": true,
	"string": true, "textwrap": true, "struct": true,
	"decimal": true, "fractions": true, "statistics": true,
	"tempfile": true, "zipfile": true, "tarfile": true,
	"gzip": true, "bz2": true, "lzma": true,
	"pickle": true, "shelve": true, "dbm": true,
	"platform": true, "ctypes": true, "multiprocessing": true,
	"concurrent": true, "asyncio": true, "signal": true,
	"traceback": true, "warnings": true, "inspect": true,
	"configparser": true, "webbrowser": true, "enum": true,
	"dataclasses": true, "contextlib": true, "operator": true,
}

var importRe = regexp.MustCompile(`(?m)^(?:import|from)\s+(\w+)`)

type Runner struct {
	dir       string
	venvDir   string
	python    string
	installed map[string]bool
}

func NewRunner(dir, pythonCmd string) *Runner {
	if dir == "" {
		dir = defaultScratchDir()
	}
	if pythonCmd == "" {
		pythonCmd = findPython()
	}
	return &Runner{
		dir:       dir,
		venvDir:   filepath.Join(dir, "venv"),
		python:    pythonCmd,
		installed: make(map[string]bool),
	}
}

func defaultScratchDir() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(os.Getenv("APPDATA"), "nsh", "scratch")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		home = os.TempDir()
	}
	return filepath.Join(home, ".config", "nsh", "scratch")
}

func findPython() string {
	for _, name := range []string{"python3", "python"} {
		if _, err := exec.LookPath(name); err == nil {
			return name
		}
	}
	return ""
}

func (r *Runner) Available() bool {
	return r.python != ""
}

func LooksLikePython(input string) bool {
	lower := strings.ToLower(input)
	for _, signal := range pythonSignals {
		if strings.Contains(lower, signal) {
			return true
		}
	}
	return false
}

func (r *Runner) DetectImports(script string) []string {
	matches := importRe.FindAllStringSubmatch(script, -1)
	seen := make(map[string]bool)
	var imports []string
	for _, m := range matches {
		mod := m[1]
		if seen[mod] || stdlibModules[mod] {
			continue
		}
		seen[mod] = true
		imports = append(imports, mod)
	}
	return imports
}

func (r *Runner) MapToPip(imports []string) []string {
	seen := make(map[string]bool)
	var packages []string
	for _, imp := range imports {
		pip, ok := importToPip[imp]
		if !ok {
			pip = imp
		}
		if !seen[pip] {
			seen[pip] = true
			packages = append(packages, pip)
		}
	}
	sort.Strings(packages)
	return packages
}

func (r *Runner) EnsureVenv() error {
	pyPath := r.venvPython()
	if _, err := os.Stat(pyPath); err == nil {
		return nil
	}

	fmt.Println("[nsh] Creating Python environment...")
	if err := os.MkdirAll(filepath.Dir(r.venvDir), 0755); err != nil {
		return err
	}

	cmd := exec.Command(r.python, "-m", "venv", r.venvDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (r *Runner) InstallMissing(packages []string) error {
	if len(packages) == 0 {
		return nil
	}

	var missing []string
	for _, pkg := range packages {
		if !r.installed[strings.ToLower(pkg)] {
			missing = append(missing, pkg)
		}
	}
	if len(missing) == 0 {
		return nil
	}

	r.refreshInstalled()

	var toInstall []string
	for _, pkg := range missing {
		if !r.installed[strings.ToLower(pkg)] {
			toInstall = append(toInstall, pkg)
		}
	}
	if len(toInstall) == 0 {
		return nil
	}

	fmt.Printf("\n[nsh] The script requires missing pip packages: %s\n", strings.Join(toInstall, ", "))
	fmt.Print("[nsh] Auto-install them? [Y/n] ")
	var confirm string
	fmt.Scanln(&confirm)
	confirm = strings.ToLower(strings.TrimSpace(confirm))
	if confirm != "" && confirm != "y" && confirm != "yes" {
		return fmt.Errorf("pip install cancelled by user")
	}

	fmt.Printf("[nsh] Installing: %s\n", strings.Join(toInstall, ", "))
	args := append([]string{"-m", "pip", "install", "-q"}, toInstall...)
	cmd := exec.Command(r.venvPython(), args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err == nil {
		for _, pkg := range toInstall {
			r.installed[strings.ToLower(pkg)] = true
		}
	}
	return err
}

func (r *Runner) Run(input, script string) error {
	if err := os.MkdirAll(r.dir, 0755); err != nil {
		return err
	}

	timestamp := time.Now().Format("2006-01-02_150405")
	slug := slugify(input)
	filename := fmt.Sprintf("%s_%s.py", timestamp, slug)
	scriptPath := filepath.Join(r.dir, filename)

	if err := os.WriteFile(scriptPath, []byte(script), 0644); err != nil {
		return err
	}

	imports := r.DetectImports(script)
	packages := r.MapToPip(imports)

	pythonExe := r.python
	if len(packages) > 0 {
		if err := r.EnsureVenv(); err != nil {
			return fmt.Errorf("venv creation failed: %w", err)
		}
		if err := r.InstallMissing(packages); err != nil {
			return fmt.Errorf("pip install failed: %w", err)
		}
		pythonExe = r.venvPython()
	}

	fmt.Printf("[nsh] Script saved: %s\n", scriptPath)
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	cmd := exec.CommandContext(ctx, pythonExe, scriptPath)
	cmd.Dir, _ = os.Getwd()
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return fmt.Errorf("script execution timed out after 10 minutes")
	}
	return err
}

func (r *Runner) refreshInstalled() {
	cmd := exec.Command(r.venvPython(), "-m", "pip", "list", "--format=columns")
	output, err := cmd.Output()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) >= 1 {
			r.installed[strings.ToLower(fields[0])] = true
		}
	}
}


func (r *Runner) venvPython() string {
	if runtime.GOOS == "windows" {
		return filepath.Join(r.venvDir, "Scripts", "python.exe")
	}
	return filepath.Join(r.venvDir, "bin", "python")
}

func (r *Runner) ScriptsDir() string {
	return r.dir
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, s)
	for strings.Contains(s, "__") {
		s = strings.ReplaceAll(s, "__", "_")
	}
	s = strings.Trim(s, "_")
	if len(s) > 40 {
		s = s[:40]
	}
	return s
}
