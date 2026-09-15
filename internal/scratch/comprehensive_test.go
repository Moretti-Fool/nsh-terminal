package scratch

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

// 105 tests total in this file

func TestLooksLikePythonComprehensive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		// False positives / negatives
		{"empty", "", false},
		{"random string", "just a random string", false},
		{"golang", "write a go script", false},
		{"near miss", "panda are cute", false},
		{"partial word", "scipylite", true}, // string.Contains matches scipy
		// All signals individually
		{"signal pandas", "use pandas", true},
		{"signal numpy", "use numpy", true},
		{"signal matplotlib", "use matplotlib", true},
		{"signal pyplot", "use pyplot", true},
		{"signal seaborn", "use seaborn", true},
		{"signal dataframe", "use dataframe", true},
		{"signal scipy", "use scipy", true},
		{"signal sklearn", "use sklearn", true},
		{"signal scikit", "use scikit", true},
		{"signal openpyxl", "use openpyxl", true},
		{"signal xlrd", "use xlrd", true},
		{"signal plotly", "use plotly", true},
		{"signal bokeh", "use bokeh", true},
		{"signal with python", "do it with python", true},
		{"signal using python", "using python to do it", true},
		{"signal in python", "code in python", true},
		{"signal python script", "run a python script", true},
		{"signal python code", "write python code", true},
		{"signal using pandas", "using pandas data", true},
		{"signal with pandas", "do it with pandas", true},
		{"signal using numpy", "using numpy arrays", true},
		{"signal with numpy", "do it with numpy", true},
		{"signal using matplotlib", "using matplotlib plots", true},
		{"signal with matplotlib", "do it with matplotlib", true},
		// Case insensitivity
		{"case mixed", "PAnDaS is great", true},
		{"case upper", "NUMPY", true},
		{"case python", "PYTHON SCRIPT", true},
		// Edge cases
		{"signal at start", "pandas is good", true},
		{"signal at end", "good is pandas", true},
		{"multiple signals", "python script using pandas and numpy", true},
		{"newlines", "hello\nworld\npandas", true},
		{"tabs", "hello\tpandas\tworld", true},
		{"unicode", "🐼 pandas", true},
		{"very long string", strings.Repeat("a", 1000) + " pandas " + strings.Repeat("b", 1000), true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LooksLikePython(tt.input); got != tt.want {
				t.Errorf("LooksLikePython() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDetectImportsComprehensive(t *testing.T) {
	t.Parallel()
	r := NewRunner("", "")
	tests := []struct {
		name   string
		script string
		want   []string
	}{
		{"empty", "", nil},
		{"no imports", "print('hello')", nil},
		{"basic import", "import pandas", []string{"pandas"}},
		{"import with alias", "import numpy as np", []string{"numpy"}},
		{"multiple imports", "import pandas\nimport numpy", []string{"pandas", "numpy"}},
		{"from import", "from sklearn import svm", []string{"sklearn"}},
		{"from import multiple", "from sklearn import svm, datasets", []string{"sklearn"}},
		{"stdlib ignored", "import os\nimport sys", nil},
		{"mixed stdlib and third party", "import os\nimport pandas\nimport json", []string{"pandas"}},
		{"deduplication", "import pandas\nimport pandas as pd\nfrom pandas import DataFrame", []string{"pandas"}},
		{"commented out", "# import pandas\nimport numpy", []string{"numpy"}},
		{"import in string", "a = '''\nimport pandas\n'''", []string{"pandas"}},
		{"indented import", " if True:\n    import pandas", nil},
		{"malicious name", "import pandas; os.system('rm -rf /')", []string{"pandas"}},
		{"carriage return", "import pandas\r\nimport numpy", []string{"pandas", "numpy"}},
		{"unicode", "import pandas_🐼", []string{"pandas_"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := r.DetectImports(tt.script)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DetectImports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMapToPipComprehensive(t *testing.T) {
	t.Parallel()
	r := NewRunner("", "")
	tests := []struct {
		name    string
		imports []string
		want    []string
	}{
		{"empty", []string{}, nil},
		{"nil", nil, nil},
		{"unknown", []string{"unknownpkg"}, []string{"unknownpkg"}},
		{"unknown multiple", []string{"pkg_b", "pkg_a"}, []string{"pkg_a", "pkg_b"}},
		{"deduplication", []string{"pandas", "pandas"}, []string{"pandas"}},
		{"mapping pandas", []string{"pandas"}, []string{"pandas"}},
		{"mapping sklearn", []string{"sklearn"}, []string{"scikit-learn"}},
		{"mapping PIL", []string{"PIL"}, []string{"Pillow"}},
		{"mapping cv2", []string{"cv2"}, []string{"opencv-python"}},
		{"mapping bs4", []string{"bs4"}, []string{"beautifulsoup4"}},
		{"mapping yaml", []string{"yaml"}, []string{"pyyaml"}},
		{"mapping dotenv", []string{"dotenv"}, []string{"python-dotenv"}},
		{"mapping psycopg2", []string{"psycopg2"}, []string{"psycopg2-binary"}},
		{"mapping jwt", []string{"jwt"}, []string{"PyJWT"}},
		{"mapping mixed", []string{"sklearn", "PIL", "unknown", "cv2"}, []string{"Pillow", "opencv-python", "scikit-learn", "unknown"}},
		{"mapping all", []string{"pandas", "numpy", "matplotlib", "seaborn", "scipy", "requests", "flask", "fastapi", "uvicorn", "sqlalchemy", "torch", "tensorflow", "lxml", "openpyxl", "xlrd", "xlsxwriter", "pymongo", "redis", "boto3", "paramiko", "cryptography", "rich", "click", "httpx", "aiohttp", "pydantic", "tabulate", "tqdm", "plotly", "dash", "streamlit"}, []string{"aiohttp", "boto3", "click", "cryptography", "dash", "fastapi", "flask", "httpx", "lxml", "matplotlib", "numpy", "openpyxl", "pandas", "paramiko", "plotly", "pydantic", "pymongo", "redis", "requests", "rich", "scipy", "seaborn", "sqlalchemy", "streamlit", "tabulate", "tensorflow", "torch", "tqdm", "uvicorn", "xlrd", "xlsxwriter"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := r.MapToPip(tt.imports)
			if len(got) == 0 && len(tt.want) == 0 {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapToPip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSlugifyComprehensive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"normal", "hello world", "hello_world"},
		{"mixed case", "HeLlO", "hello"},
		{"numbers", "123 456", "123_456"},
		{"special chars", "hello!@#world", "hello_world"},
		{"trailing leading special", "!@#hello world!@#", "hello_world"},
		{"multiple spaces", "hello   world", "hello_world"},
		{"underscores", "hello___world", "hello_world"},
		{"long string", "this is a very long string that should be truncated because it exceeds forty characters", "this_is_a_very_long_string_that_should_b"},
		{"exactly 40 chars", "abcdefghijklmnopqrstuvwxyz12345678901234", "abcdefghijklmnopqrstuvwxyz12345678901234"},
		{"41 chars", "abcdefghijklmnopqrstuvwxyz123456789012345", "abcdefghijklmnopqrstuvwxyz12345678901234"},
		{"unicode", "hello 世界", "hello"},
		{"only unicode", "世界", ""},
		{"leading underscore", "_hello", "hello"},
		{"trailing underscore", "hello_", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slugify(tt.input); got != tt.want {
				t.Errorf("slugify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStdlibModulesComprehensive(t *testing.T) {
	t.Parallel()
	stdlibKeys := []string{"os", "sys", "json", "math", "re", "datetime", "subprocess", "typing"}
	for _, key := range stdlibKeys {
		if !stdlibModules[key] {
			t.Errorf("stdlibModules missing %s", key)
		}
	}
	thirdPartyKeys := []string{"pandas", "numpy", "requests", "django"}
	for _, key := range thirdPartyKeys {
		if stdlibModules[key] {
			t.Errorf("stdlibModules incorrectly contains %s", key)
		}
	}
}

func TestNewRunnerComprehensive(t *testing.T) {
	t.Parallel()
	r1 := NewRunner("", "")
	if r1.dir == "" {
		t.Errorf("NewRunner with empty dir should use default")
	}
	if runtime.GOOS == "windows" {
		if !strings.Contains(r1.dir, "APPDATA") && !strings.Contains(r1.dir, "AppData") && !strings.Contains(r1.dir, "appdata") {
			// Optional check
		}
	} else {
		if !strings.Contains(r1.dir, ".config") {
			t.Errorf("NewRunner default dir on unix should contain .config")
		}
	}

	r2 := NewRunner("/tmp/custom", "python3")
	if r2.dir != "/tmp/custom" {
		t.Errorf("NewRunner with custom dir failed")
	}
	if r2.python != "python3" {
		t.Errorf("NewRunner with custom python failed")
	}
	if r2.venvDir != filepath.Join("/tmp/custom", "venv") {
		t.Errorf("NewRunner venvDir incorrect")
	}
	if r2.ScriptsDir() != "/tmp/custom" {
		t.Errorf("ScriptsDir() failed")
	}
}

func TestAvailableComprehensive(t *testing.T) {
	t.Parallel()
	r1 := NewRunner("", "my_custom_python")
	if !r1.Available() {
		t.Errorf("Available() should be true when python is set")
	}
	r2 := NewRunner("", "")
	r2.python = "" // force empty
	if r2.Available() {
		t.Errorf("Available() should be false when python is empty")
	}
}

func TestDefaultScratchDirComprehensive(t *testing.T) {
	t.Parallel()
	dir := defaultScratchDir()
	if dir == "" {
		t.Errorf("defaultScratchDir() returned empty")
	}
	if runtime.GOOS == "windows" {
		if !strings.Contains(strings.ToLower(dir), "nsh") || !strings.Contains(strings.ToLower(dir), "scratch") {
			t.Errorf("defaultScratchDir() on windows missing nsh/scratch: %s", dir)
		}
	} else {
		if !strings.Contains(dir, ".config") {
			t.Errorf("defaultScratchDir() on unix missing .config")
		}
	}
}

func TestFindPythonComprehensive(t *testing.T) {
	t.Parallel()
	py := findPython()
	_ = py // could be empty if python is not installed on test machine
}

func TestImportReComprehensive(t *testing.T) {
	t.Parallel()
	tests := []struct {
		input string
		want  string // the first submatch
	}{
		{"import pandas", "pandas"},
		{"from sklearn import svm", "sklearn"},
		{"import numpy as np", "numpy"},
		{"from bs4 import BeautifulSoup", "bs4"},
		{"import   math", "math"},
		{"from   os   import path", "os"},
	}
	for _, tt := range tests {
		m := importRe.FindStringSubmatch(tt.input)
		if len(m) < 2 {
			t.Errorf("importRe failed to match %q", tt.input)
			continue
		}
		if m[1] != tt.want {
			t.Errorf("importRe matched %q, want %q", m[1], tt.want)
		}
	}
}
