//go:build windows

package repl

import (
	"bytes"
	"testing"
)

func newTestReader() *consoleReader {
	return &consoleReader{}
}

func sendKeyEvent(r *consoleReader, bKeyDown int32, vCode uint16, uChar uint16, ctrlState uint32, repeat uint16) {
	if repeat == 0 {
		repeat = 1
	}
	ker := &keyEventRecord{
		bKeyDown:          bKeyDown,
		wRepeatCount:      repeat,
		wVirtualKeyCode:   vCode,
		unicodeChar:       uChar,
		dwControlKeyState: ctrlState,
	}
	if bKeyDown != 0 {
		r.processKeyEvent(ker)
	}
}

func getBuf(r *consoleReader) []byte {
	out := make([]byte, len(r.buf))
	copy(out, r.buf)
	r.buf = r.buf[:0]
	return out
}

func TestConsoleReader_NormalKeys(t *testing.T) {
	r := newTestReader()

	// Type 'd', 'f', 'b' normally
	sendKeyEvent(r, 1, 'D', 'd', 0, 1)
	sendKeyEvent(r, 1, 'F', 'f', 0, 1)
	sendKeyEvent(r, 1, 'B', 'b', 0, 1)

	got := getBuf(r)
	if string(got) != "dfb" {
		t.Fatalf("expected 'dfb', got %q", string(got))
	}
}

func TestConsoleReader_Backspace(t *testing.T) {
	r := newTestReader()

	sendKeyEvent(r, 1, vkBack, '\b', 0, 1)
	got := getBuf(r)
	if len(got) != 1 || got[0] != '\b' {
		t.Fatalf("expected single \\b (8), got %v", got)
	}
}

func TestConsoleReader_NoPhantomAltAfterFocusChange(t *testing.T) {
	r := newTestReader()

	// Simulate user pressing Alt+Tab:
	// Alt KeyDown
	sendKeyEvent(r, 1, vkMenu, 0, leftAltPressed, 1)
	// (KeyUp is lost because window lost focus)

	// User switches back and types 'd', 'f', 'b' without Alt
	sendKeyEvent(r, 1, 'D', 'd', 0, 1)
	sendKeyEvent(r, 1, 'F', 'f', 0, 1)
	sendKeyEvent(r, 1, 'B', 'b', 0, 1)
	// User presses Backspace
	sendKeyEvent(r, 1, vkBack, '\b', 0, 1)

	got := getBuf(r)
	expected := "dfb\b"
	if string(got) != expected {
		t.Fatalf("expected %q, got %q (hex: %x)", expected, string(got), got)
	}
}

func TestConsoleReader_AltKeys(t *testing.T) {
	r := newTestReader()

	// Alt+b -> \033b
	sendKeyEvent(r, 1, 'B', 'b', leftAltPressed, 1)
	if got := string(getBuf(r)); got != "\033b" {
		t.Fatalf("expected \\033b, got %q", got)
	}

	// Alt+f -> \033f
	sendKeyEvent(r, 1, 'F', 'f', rightAltPressed, 1)
	if got := string(getBuf(r)); got != "\033f" {
		t.Fatalf("expected \\033f, got %q", got)
	}

	// Alt+d -> \033d
	sendKeyEvent(r, 1, 'D', 'd', leftAltPressed, 1)
	if got := string(getBuf(r)); got != "\033d" {
		t.Fatalf("expected \\033d, got %q", got)
	}

	// Alt+Backspace -> \033\x7f
	sendKeyEvent(r, 1, vkBack, '\b', leftAltPressed, 1)
	if got := getBuf(r); !bytes.Equal(got, []byte{'\033', charBackspace}) {
		t.Fatalf("expected \\033\\x7f, got %v", got)
	}
}

func TestConsoleReader_CtrlKeys(t *testing.T) {
	r := newTestReader()

	// Ctrl+Backspace -> charCtrlW (\x17 = 23)
	sendKeyEvent(r, 1, vkBack, '\b', leftCtrlPressed, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charCtrlW {
		t.Fatalf("expected charCtrlW (23), got %v", got)
	}

	// Ctrl+Left -> \033b
	sendKeyEvent(r, 1, vkLeft, 0, leftCtrlPressed, 1)
	if got := string(getBuf(r)); got != "\033b" {
		t.Fatalf("expected \\033b, got %q", got)
	}

	// Ctrl+Right -> \033f
	sendKeyEvent(r, 1, vkRight, 0, rightCtrlPressed, 1)
	if got := string(getBuf(r)); got != "\033f" {
		t.Fatalf("expected \\033f, got %q", got)
	}

	// Ctrl+Delete -> \033d
	sendKeyEvent(r, 1, vkDelete, 0, leftCtrlPressed, 1)
	if got := string(getBuf(r)); got != "\033d" {
		t.Fatalf("expected \\033d, got %q", got)
	}

	// Ctrl+C -> 3
	sendKeyEvent(r, 1, 'C', 3, leftCtrlPressed, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charInterrupt {
		t.Fatalf("expected charInterrupt (3), got %v", got)
	}

	// Ctrl+D -> 4
	sendKeyEvent(r, 1, 'D', 4, leftCtrlPressed, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charDelete {
		t.Fatalf("expected charDelete (4), got %v", got)
	}
}

func TestConsoleReader_NavigationKeys(t *testing.T) {
	r := newTestReader()

	sendKeyEvent(r, 1, vkLeft, 0, 0, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charBackward {
		t.Fatalf("expected charBackward (2), got %v", got)
	}

	sendKeyEvent(r, 1, vkRight, 0, 0, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charForward {
		t.Fatalf("expected charForward (6), got %v", got)
	}

	sendKeyEvent(r, 1, vkUp, 0, 0, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charPrev {
		t.Fatalf("expected charPrev (16), got %v", got)
	}

	sendKeyEvent(r, 1, vkDown, 0, 0, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charNext {
		t.Fatalf("expected charNext (14), got %v", got)
	}

	sendKeyEvent(r, 1, vkHome, 0, 0, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charLineStart {
		t.Fatalf("expected charLineStart (1), got %v", got)
	}

	sendKeyEvent(r, 1, vkEnd, 0, 0, 1)
	if got := getBuf(r); len(got) != 1 || got[0] != charLineEnd {
		t.Fatalf("expected charLineEnd (5), got %v", got)
	}

	sendKeyEvent(r, 1, vkDelete, 0, 0, 1)
	if got := string(getBuf(r)); got != "\033[3~" {
		t.Fatalf("expected \\033[3~, got %q", got)
	}

	sendKeyEvent(r, 1, vkTab, '\t', 0, 1)
	if got := string(getBuf(r)); got != "\t" {
		t.Fatalf("expected \\t, got %q", got)
	}

	sendKeyEvent(r, 1, vkTab, '\t', shiftPressed, 1)
	if got := string(getBuf(r)); got != "\033[Z" {
		t.Fatalf("expected \\033[Z, got %q", got)
	}
}

func TestConsoleReader_AltGr(t *testing.T) {
	r := newTestReader()

	// AltGr sends both rightAltPressed and leftCtrlPressed
	altGrState := uint32(rightAltPressed | leftCtrlPressed)
	sendKeyEvent(r, 1, '2', '@', altGrState, 1)
	got := string(getBuf(r))
	if got != "@" {
		t.Fatalf("expected '@' without escape, got %q", got)
	}
}

func TestConsoleReader_Surrogates(t *testing.T) {
	r := newTestReader()

	// Rocket emoji: U+1F680 -> UTF-16: 0xD83D 0xDE80
	sendKeyEvent(r, 1, 0, 0xD83D, 0, 1)
	if len(r.buf) != 0 {
		t.Fatalf("expected surrogate high to be buffered, got %v", r.buf)
	}
	sendKeyEvent(r, 1, 0, 0xDE80, 0, 1)
	got := string(getBuf(r))
	if got != "🚀" {
		t.Fatalf("expected rocket emoji, got %q", got)
	}
}

func TestConsoleReader_StandaloneModifiersIgnored(t *testing.T) {
	r := newTestReader()

	sendKeyEvent(r, 1, vkShift, 0, shiftPressed, 1)
	sendKeyEvent(r, 1, vkControl, 0, leftCtrlPressed, 1)
	sendKeyEvent(r, 1, vkMenu, 0, leftAltPressed, 1)
	sendKeyEvent(r, 1, vkCapital, 0, 0, 1)

	got := getBuf(r)
	if len(got) != 0 {
		t.Fatalf("expected 0 bytes for standalone modifiers, got %v", got)
	}
}

func TestConsoleReader_KeyUpIgnored(t *testing.T) {
	r := newTestReader()

	// KeyUp events (bKeyDown == 0)
	sendKeyEvent(r, 0, 'A', 'a', 0, 1)
	got := getBuf(r)
	if len(got) != 0 {
		t.Fatalf("expected 0 bytes for key up, got %v", got)
	}
}

func TestConsoleReader_RepeatCount(t *testing.T) {
	r := newTestReader()

	sendKeyEvent(r, 1, 'X', 'x', 0, 3)
	got := string(getBuf(r))
	if got != "xxx" {
		t.Fatalf("expected 'xxx', got %q", got)
	}
}
