//go:build !windows

package repl

import "io"

func newConsoleReader() io.ReadCloser {
	return nil
}
