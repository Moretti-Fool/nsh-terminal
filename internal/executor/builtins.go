package executor

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type lsFlags struct {
	long    bool
	all     bool
	sortMod bool
	reverse bool
	human   bool
}

var lsFlagChars = map[byte]bool{'l': true, 'a': true, 't': true, 'r': true, 'h': true}

func looksLikeFlags(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !lsFlagChars[s[i]] {
			return false
		}
	}
	return true
}

func parseLsFlags(tokens []string) (lsFlags, []string) {
	var f lsFlags
	var rest []string
	for _, t := range tokens {
		flagStr := ""
		if strings.HasPrefix(t, "-") && len(t) > 1 && t[1] != '-' {
			flagStr = t[1:]
		} else if looksLikeFlags(t) {
			flagStr = t
		}

		if flagStr != "" {
			for _, ch := range flagStr {
				switch ch {
				case 'l':
					f.long = true
				case 'a':
					f.all = true
				case 't':
					f.sortMod = true
				case 'r':
					f.reverse = true
				case 'h':
					f.human = true
				}
			}
		} else {
			rest = append(rest, t)
		}
	}
	return f, rest
}

func builtinLsFull(tokens []string) (RunResult, bool) {
	start := time.Now()
	flags, paths := parseLsFlags(tokens[1:])
	if len(paths) == 0 {
		paths = []string{"."}
	}

	var buf strings.Builder
	for _, dir := range paths {
		if len(paths) > 1 {
			buf.WriteString(dir + ":\n")
		}
		entries, err := os.ReadDir(dir)
		if err != nil {
			errMsg := fmt.Sprintf("ls: %s: %v\n", dir, err)
			fmt.Fprint(os.Stderr, errMsg)
			return RunResult{Output: errMsg, ExitCode: 1, DurationMs: time.Since(start).Milliseconds()}, true
		}

		type fileEntry struct {
			name    string
			size    int64
			modTime time.Time
			isDir   bool
			mode    os.FileMode
		}

		var items []fileEntry
		for _, e := range entries {
			name := e.Name()
			if !flags.all && strings.HasPrefix(name, ".") {
				continue
			}
			info, err := e.Info()
			if err != nil {
				continue
			}
			items = append(items, fileEntry{
				name:    name,
				size:    info.Size(),
				modTime: info.ModTime(),
				isDir:   e.IsDir(),
				mode:    info.Mode(),
			})
		}

		if flags.sortMod {
			sort.Slice(items, func(i, j int) bool {
				return items[i].modTime.Before(items[j].modTime)
			})
		} else {
			sort.Slice(items, func(i, j int) bool {
				return strings.ToLower(items[i].name) < strings.ToLower(items[j].name)
			})
		}

		if flags.reverse {
			for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
				items[i], items[j] = items[j], items[i]
			}
		}

		if flags.long {
			for _, item := range items {
				typeChar := "-"
				if item.isDir {
					typeChar = "d"
				}
				sizeStr := fmt.Sprintf("%d", item.size)
				if flags.human {
					sizeStr = humanSize(item.size)
				}
				name := item.name
				if item.isDir {
					name += "/"
				}
				line := fmt.Sprintf("%srw-  %8s  %s  %s\n",
					typeChar,
					sizeStr,
					item.modTime.Format("Jan 02 15:04"),
					name)
				buf.WriteString(line)
			}
		} else {
			var names []string
			for _, item := range items {
				name := item.name
				if item.isDir {
					name += "/"
				}
				names = append(names, name)
			}
			buf.WriteString(formatColumns(names))
		}

		if len(paths) > 1 {
			buf.WriteByte('\n')
		}
	}

	output := buf.String()
	fmt.Print(output)
	return RunResult{Output: output, DurationMs: time.Since(start).Milliseconds()}, true
}

func humanSize(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(b)/float64(div), "KMGTPE"[exp])
}

func formatColumns(names []string) string {
	if len(names) == 0 {
		return ""
	}
	maxLen := 0
	for _, n := range names {
		if len(n) > maxLen {
			maxLen = len(n)
		}
	}
	colWidth := maxLen + 2
	cols := 80 / colWidth
	if cols < 1 {
		cols = 1
	}
	var buf strings.Builder
	for i, name := range names {
		if i > 0 && i%cols == 0 {
			buf.WriteByte('\n')
		}
		buf.WriteString(fmt.Sprintf("%-*s", colWidth, name))
	}
	buf.WriteByte('\n')
	return buf.String()
}

func builtinCat(tokens []string) (RunResult, bool) {
	start := time.Now()
	if len(tokens) < 2 {
		return RunResult{Output: "cat: missing file\n", ExitCode: 1, DurationMs: 0}, true
	}

	var buf strings.Builder
	exitCode := 0
	for _, path := range tokens[1:] {
		if strings.HasPrefix(path, "-") {
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			msg := fmt.Sprintf("cat: %s: %v\n", path, err)
			fmt.Fprint(os.Stderr, msg)
			buf.WriteString(msg)
			exitCode = 1
			continue
		}
		content := string(data)
		fmt.Print(content)
		buf.WriteString(content)
	}
	return RunResult{Output: buf.String(), ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinHead(tokens []string) (RunResult, bool) {
	start := time.Now()
	n := 10
	var files []string
	for i := 1; i < len(tokens); i++ {
		if tokens[i] == "-n" && i+1 < len(tokens) {
			fmt.Sscanf(tokens[i+1], "%d", &n)
			i++
		} else if !strings.HasPrefix(tokens[i], "-") {
			files = append(files, tokens[i])
		}
	}
	if len(files) == 0 {
		return RunResult{Output: "head: missing file\n", ExitCode: 1}, true
	}

	var buf strings.Builder
	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			msg := fmt.Sprintf("head: %s: %v\n", path, err)
			fmt.Fprint(os.Stderr, msg)
			buf.WriteString(msg)
			continue
		}
		if len(files) > 1 {
			header := fmt.Sprintf("==> %s <==\n", path)
			fmt.Print(header)
			buf.WriteString(header)
		}
		scanner := bufio.NewScanner(f)
		for i := 0; i < n && scanner.Scan(); i++ {
			line := scanner.Text() + "\n"
			fmt.Print(line)
			buf.WriteString(line)
		}
		f.Close()
	}
	return RunResult{Output: buf.String(), DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinTail(tokens []string) (RunResult, bool) {
	start := time.Now()
	n := 10
	var files []string
	for i := 1; i < len(tokens); i++ {
		if tokens[i] == "-n" && i+1 < len(tokens) {
			fmt.Sscanf(tokens[i+1], "%d", &n)
			i++
		} else if !strings.HasPrefix(tokens[i], "-") {
			files = append(files, tokens[i])
		}
	}
	if len(files) == 0 {
		return RunResult{Output: "tail: missing file\n", ExitCode: 1}, true
	}

	var buf strings.Builder
	for _, path := range files {
		data, err := os.ReadFile(path)
		if err != nil {
			msg := fmt.Sprintf("tail: %s: %v\n", path, err)
			fmt.Fprint(os.Stderr, msg)
			buf.WriteString(msg)
			continue
		}
		if len(files) > 1 {
			header := fmt.Sprintf("==> %s <==\n", path)
			fmt.Print(header)
			buf.WriteString(header)
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 && lines[len(lines)-1] == "" {
			lines = lines[:len(lines)-1]
		}
		startIdx := len(lines) - n
		if startIdx < 0 {
			startIdx = 0
		}
		for _, line := range lines[startIdx:] {
			out := line + "\n"
			fmt.Print(out)
			buf.WriteString(out)
		}
		f, _ := os.Open(path)
		if f != nil {
			f.Close()
		}
	}
	return RunResult{Output: buf.String(), DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinWc(tokens []string) (RunResult, bool) {
	start := time.Now()
	if len(tokens) < 2 {
		return RunResult{Output: "wc: missing file\n", ExitCode: 1}, true
	}

	var buf strings.Builder
	for _, path := range tokens[1:] {
		if strings.HasPrefix(path, "-") {
			continue
		}
		data, err := os.ReadFile(path)
		if err != nil {
			msg := fmt.Sprintf("wc: %s: %v\n", path, err)
			fmt.Fprint(os.Stderr, msg)
			buf.WriteString(msg)
			continue
		}
		content := string(data)
		lines := strings.Count(content, "\n")
		words := len(strings.Fields(content))
		chars := len(data)
		line := fmt.Sprintf("%8d %8d %8d %s\n", lines, words, chars, path)
		fmt.Print(line)
		buf.WriteString(line)
	}
	return RunResult{Output: buf.String(), DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinTouch(tokens []string) (RunResult, bool) {
	start := time.Now()
	if len(tokens) < 2 {
		return RunResult{Output: "touch: missing file\n", ExitCode: 1}, true
	}
	exitCode := 0
	for _, path := range tokens[1:] {
		if strings.HasPrefix(path, "-") {
			continue
		}
		if _, err := os.Stat(path); os.IsNotExist(err) {
			f, err := os.Create(path)
			if err != nil {
				fmt.Fprintf(os.Stderr, "touch: %v\n", err)
				exitCode = 1
				continue
			}
			f.Close()
		} else {
			now := time.Now()
			os.Chtimes(path, now, now)
		}
	}
	return RunResult{ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinMkdir(tokens []string) (RunResult, bool) {
	start := time.Now()
	parents := false
	var dirs []string
	for _, t := range tokens[1:] {
		if t == "-p" {
			parents = true
		} else if !strings.HasPrefix(t, "-") {
			dirs = append(dirs, t)
		}
	}
	if len(dirs) == 0 {
		return RunResult{Output: "mkdir: missing directory\n", ExitCode: 1}, true
	}
	exitCode := 0
	for _, dir := range dirs {
		var err error
		if parents {
			err = os.MkdirAll(dir, 0755)
		} else {
			err = os.Mkdir(dir, 0755)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "mkdir: %v\n", err)
			exitCode = 1
		}
	}
	return RunResult{ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinCp(tokens []string) (RunResult, bool) {
	start := time.Now()
	var args []string
	recursive := false
	for _, t := range tokens[1:] {
		if t == "-r" || t == "-R" || t == "--recursive" {
			recursive = true
		} else if !strings.HasPrefix(t, "-") {
			args = append(args, t)
		}
	}
	if len(args) < 2 {
		return RunResult{Output: "cp: missing source or destination\n", ExitCode: 1}, true
	}
	dst := args[len(args)-1]
	sources := args[:len(args)-1]

	exitCode := 0
	for _, src := range sources {
		info, err := os.Stat(src)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cp: %v\n", err)
			exitCode = 1
			continue
		}
		if info.IsDir() && !recursive {
			fmt.Fprintf(os.Stderr, "cp: %s is a directory (use -r)\n", src)
			exitCode = 1
			continue
		}
		if info.IsDir() {
			err = copyDir(src, filepath.Join(dst, filepath.Base(src)))
		} else {
			target := dst
			if di, e := os.Stat(dst); e == nil && di.IsDir() {
				target = filepath.Join(dst, filepath.Base(src))
			}
			err = copyFile(src, target)
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "cp: %v\n", err)
			exitCode = 1
		}
	}
	return RunResult{ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}

func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, _ := filepath.Rel(src, path)
		target := filepath.Join(dst, rel)
		if info.IsDir() {
			return os.MkdirAll(target, 0755)
		}
		return copyFile(path, target)
	})
}

func builtinMv(tokens []string) (RunResult, bool) {
	start := time.Now()
	var args []string
	for _, t := range tokens[1:] {
		if !strings.HasPrefix(t, "-") {
			args = append(args, t)
		}
	}
	if len(args) < 2 {
		return RunResult{Output: "mv: missing source or destination\n", ExitCode: 1}, true
	}
	dst := args[len(args)-1]
	sources := args[:len(args)-1]

	exitCode := 0
	for _, src := range sources {
		target := dst
		if di, e := os.Stat(dst); e == nil && di.IsDir() {
			target = filepath.Join(dst, filepath.Base(src))
		}
		if err := os.Rename(src, target); err != nil {
			fmt.Fprintf(os.Stderr, "mv: %v\n", err)
			exitCode = 1
		}
	}
	return RunResult{ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinRm(tokens []string) (RunResult, bool) {
	start := time.Now()
	recursive := false
	force := false
	var paths []string
	for _, t := range tokens[1:] {
		if strings.HasPrefix(t, "-") && !strings.HasPrefix(t, "--") {
			for _, ch := range t[1:] {
				switch ch {
				case 'r', 'R':
					recursive = true
				case 'f':
					force = true
				}
			}
		} else if !strings.HasPrefix(t, "-") {
			paths = append(paths, t)
		}
	}
	if len(paths) == 0 {
		return RunResult{Output: "rm: missing file\n", ExitCode: 1}, true
	}
	exitCode := 0
	for _, p := range paths {
		info, err := os.Stat(p)
		if err != nil {
			if !force {
				fmt.Fprintf(os.Stderr, "rm: %v\n", err)
				exitCode = 1
			}
			continue
		}
		if info.IsDir() && !recursive {
			fmt.Fprintf(os.Stderr, "rm: %s: is a directory (use -r)\n", p)
			exitCode = 1
			continue
		}
		if recursive {
			err = os.RemoveAll(p)
		} else {
			err = os.Remove(p)
		}
		if err != nil && !force {
			fmt.Fprintf(os.Stderr, "rm: %v\n", err)
			exitCode = 1
		}
	}
	return RunResult{ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinEcho(tokens []string) (RunResult, bool) {
	start := time.Now()
	output := strings.Join(tokens[1:], " ") + "\n"
	fmt.Print(output)
	return RunResult{Output: output, DurationMs: time.Since(start).Milliseconds()}, true
}

func builtinGrep(tokens []string) (RunResult, bool) {
	start := time.Now()
	if len(tokens) < 3 {
		return RunResult{Output: "grep: usage: grep PATTERN FILE...\n", ExitCode: 1}, true
	}

	ignoreCase := false
	lineNum := false
	var pattern string
	var files []string
	patternSet := false
	for i := 1; i < len(tokens); i++ {
		t := tokens[i]
		if strings.HasPrefix(t, "-") && !patternSet {
			for _, ch := range t[1:] {
				switch ch {
				case 'i':
					ignoreCase = true
				case 'n':
					lineNum = true
				}
			}
		} else if !patternSet {
			pattern = t
			patternSet = true
		} else {
			files = append(files, t)
		}
	}
	if pattern == "" || len(files) == 0 {
		return RunResult{Output: "grep: usage: grep PATTERN FILE...\n", ExitCode: 1}, true
	}

	matchPattern := pattern
	if ignoreCase {
		matchPattern = strings.ToLower(pattern)
	}

	var buf strings.Builder
	found := false
	multiFile := len(files) > 1

	for _, path := range files {
		f, err := os.Open(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "grep: %s: %v\n", path, err)
			continue
		}
		scanner := bufio.NewScanner(f)
		num := 0
		for scanner.Scan() {
			num++
			line := scanner.Text()
			check := line
			if ignoreCase {
				check = strings.ToLower(line)
			}
			if strings.Contains(check, matchPattern) {
				found = true
				var out string
				prefix := ""
				if multiFile {
					prefix = path + ":"
				}
				if lineNum {
					out = fmt.Sprintf("%s%d:%s\n", prefix, num, line)
				} else {
					out = fmt.Sprintf("%s%s\n", prefix, line)
				}
				fmt.Print(out)
				buf.WriteString(out)
			}
		}
		f.Close()
	}

	exitCode := 0
	if !found {
		exitCode = 1
	}
	return RunResult{Output: buf.String(), ExitCode: exitCode, DurationMs: time.Since(start).Milliseconds()}, true
}
