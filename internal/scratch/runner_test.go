package scratch

import (
	"testing"
)

func TestLooksLikePython(t *testing.T) {
	for _, input := range []string{
		"open sales.csv with pandas",
		"plot this data using matplotlib",
		"analyze data.csv using numpy",
		"write a python script to merge files",
		"read excel file with openpyxl",
		"create a dataframe from this csv",
	} {
		if !LooksLikePython(input) {
			t.Errorf("LooksLikePython(%q) = false, want true", input)
		}
	}

	for _, input := range []string{
		"ls -la",
		"git status",
		"show me all files",
		"what is golang",
		"find . -name *.go",
	} {
		if LooksLikePython(input) {
			t.Errorf("LooksLikePython(%q) = true, want false", input)
		}
	}
}

func TestDetectImports(t *testing.T) {
	r := NewRunner("", "")
	script := `import pandas as pd
import numpy as np
from sklearn.ensemble import RandomForestClassifier
import os
import json
import matplotlib.pyplot as plt
`
	imports := r.DetectImports(script)

	expected := map[string]bool{
		"pandas":     true,
		"numpy":      true,
		"sklearn":    true,
		"matplotlib": true,
	}

	for _, imp := range imports {
		if !expected[imp] {
			t.Errorf("unexpected import: %s", imp)
		}
		delete(expected, imp)
	}
	for imp := range expected {
		t.Errorf("missing import: %s", imp)
	}
}

func TestMapToPip(t *testing.T) {
	r := NewRunner("", "")
	imports := []string{"pandas", "sklearn", "PIL", "bs4"}
	packages := r.MapToPip(imports)

	expected := map[string]bool{
		"pandas":          true,
		"scikit-learn":    true,
		"Pillow":          true,
		"beautifulsoup4":  true,
	}

	for _, pkg := range packages {
		if !expected[pkg] {
			t.Errorf("unexpected pip package: %s", pkg)
		}
		delete(expected, pkg)
	}
	for pkg := range expected {
		t.Errorf("missing pip package: %s", pkg)
	}
}

func TestSlugify(t *testing.T) {
	tests := map[string]string{
		"open sales.csv with pandas":  "open_sales_csv_with_pandas",
		"plot monthly revenue":        "plot_monthly_revenue",
		"Hello World!":                "hello_world",
	}
	for input, want := range tests {
		got := slugify(input)
		if got != want {
			t.Errorf("slugify(%q) = %q, want %q", input, got, want)
		}
	}
}
