//go:build windows

package repl

import (
	"io"
	"sync"
	"syscall"
	"unicode/utf16"
	"unicode/utf8"
	"unsafe"
)

var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procReadConsoleInputW  = kernel32.NewProc("ReadConsoleInputW")
	procWriteConsoleInputW = kernel32.NewProc("WriteConsoleInputW")
	procGetConsoleMode     = kernel32.NewProc("GetConsoleMode")
)

const (
	eventKey   = 0x0001
	eventFocus = 0x0010

	// Control key states
	rightAltPressed  = 0x0001
	leftAltPressed   = 0x0002
	rightCtrlPressed = 0x0004
	leftCtrlPressed  = 0x0008
	shiftPressed     = 0x0010
	enhancedKey      = 0x0100

	// Virtual Key Codes
	vkBack    = 0x08
	vkTab     = 0x09
	vkReturn  = 0x0D
	vkShift   = 0x10
	vkControl = 0x11
	vkMenu    = 0x12 // Alt
	vkPause   = 0x13
	vkCapital = 0x14 // Caps Lock
	vkEscape  = 0x1B
	vkSpace   = 0x20
	vkEnd     = 0x23
	vkHome    = 0x24
	vkLeft    = 0x25
	vkUp      = 0x26
	vkRight   = 0x27
	vkDown    = 0x28
	vkDelete  = 0x2E
	vkLWin    = 0x5B
	vkRWin    = 0x5C
	vkApps    = 0x5D
	vkNumLock = 0x90
	vkScroll  = 0x91
	vkLShift  = 0xA0
	vkRShift  = 0xA1
	vkLControl= 0xA2
	vkRControl= 0xA3
	vkLMenu   = 0xA4
	vkRMenu   = 0xA5

	// Readline control codes
	charLineStart = 1
	charBackward  = 2
	charInterrupt = 3
	charDelete    = 4
	charLineEnd   = 5
	charForward   = 6
	charCtrlH     = 8
	charTab       = 9
	charCtrlJ     = 10
	charKill      = 11
	charCtrlL     = 12
	charEnter     = 13
	charNext      = 14
	charPrev      = 16
	charBckSearch = 18
	charFwdSearch = 19
	charTranspose = 20
	charCtrlU     = 21
	charCtrlW     = 23
	charCtrlY     = 25
	charCtrlZ     = 26
	charEsc       = 27
	charBackspace = 127
)

type inputRecord struct {
	eventType uint16
	padding   uint16
	event     [16]byte
}

type keyEventRecord struct {
	bKeyDown          int32
	wRepeatCount      uint16
	wVirtualKeyCode   uint16
	wVirtualScanCode  uint16
	unicodeChar       uint16
	dwControlKeyState uint32
}

// consoleReader implements io.ReadCloser for Windows Console input.
// Unlike readline's default RawReader on Windows, this reader reads modifier
// key state directly from dwControlKeyState on each key event, preventing
// phantom "stuck Alt" or "stuck Ctrl" bugs when the user switches windows
// with Alt+Tab or clicks outside the terminal.
type consoleReader struct {
	handle    syscall.Handle
	buf       []byte
	surrogate rune
	mu        sync.Mutex
	closed    bool
}

func newConsoleReader() io.ReadCloser {
	handle, err := syscall.GetStdHandle(syscall.STD_INPUT_HANDLE)
	if err != nil || handle == syscall.InvalidHandle {
		return nil
	}
	var mode uint32
	r1, _, _ := procGetConsoleMode.Call(uintptr(handle), uintptr(unsafe.Pointer(&mode)))
	if r1 == 0 {
		return nil
	}
	return &consoleReader{
		handle: handle,
	}
}

func (r *consoleReader) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}

	for {
		r.mu.Lock()
		if r.closed {
			r.mu.Unlock()
			return 0, io.EOF
		}
		if len(r.buf) > 0 {
			n := copy(p, r.buf)
			r.buf = r.buf[n:]
			r.mu.Unlock()
			return n, nil
		}
		r.mu.Unlock()

		var record inputRecord
		var nRead uint32
		r1, _, err := procReadConsoleInputW.Call(
			uintptr(r.handle),
			uintptr(unsafe.Pointer(&record)),
			1,
			uintptr(unsafe.Pointer(&nRead)),
		)
		if r1 == 0 {
			r.mu.Lock()
			closed := r.closed
			r.mu.Unlock()
			if closed {
				return 0, io.EOF
			}
			return 0, err
		}
		if nRead == 0 {
			continue
		}

		r.mu.Lock()
		if r.closed {
			r.mu.Unlock()
			return 0, io.EOF
		}
		r.handleRecord(&record)
		r.mu.Unlock()
	}
}

func (r *consoleReader) handleRecord(record *inputRecord) {
	if record.eventType != eventKey {
		return
	}
	ker := (*keyEventRecord)(unsafe.Pointer(&record.event[0]))
	if ker.bKeyDown == 0 {
		// Key up event - do not emit any characters.
		return
	}
	r.processKeyEvent(ker)
}

func (r *consoleReader) processKeyEvent(ker *keyEventRecord) {
	// Standalone modifier keys produce no characters.
	switch ker.wVirtualKeyCode {
	case vkShift, vkLShift, vkRShift,
		vkControl, vkLControl, vkRControl,
		vkMenu, vkLMenu, vkRMenu,
		vkCapital, vkNumLock, vkScroll,
		vkLWin, vkRWin, vkApps, vkPause:
		return
	}

	// AltGr produces both RIGHT_ALT and LEFT_CTRL on Windows international layouts.
	// When AltGr is pressed to produce characters (e.g. @, ~, \), treat as regular character.
	isAltGr := (ker.dwControlKeyState&(rightAltPressed|leftCtrlPressed)) == (rightAltPressed|leftCtrlPressed)

	var isAlt, isCtrl bool
	if !isAltGr {
		isAlt = (ker.dwControlKeyState & (rightAltPressed | leftAltPressed)) != 0
		isCtrl = (ker.dwControlKeyState & (rightCtrlPressed | leftCtrlPressed)) != 0
	}
	isShift := (ker.dwControlKeyState & shiftPressed) != 0

	var out []byte

	if isAlt && !isCtrl {
		// Alt+Key combinations
		if ker.wVirtualKeyCode == vkBack || ker.unicodeChar == vkBack {
			// Alt+Backspace -> MetaBackspace (kill word backward)
			out = append(out, '\033', charBackspace)
		} else if ker.unicodeChar != 0 {
			out = append(out, '\033')
			var b [4]byte
			n := utf8.EncodeRune(b[:], rune(ker.unicodeChar))
			out = append(out, b[:n]...)
		} else {
			switch ker.wVirtualKeyCode {
			case vkLeft:
				out = append(out, '\033', 'b') // MetaBackward
			case vkRight:
				out = append(out, '\033', 'f') // MetaForward
			case vkDelete:
				out = append(out, '\033', 'd') // MetaDelete
			case vkBack:
				out = append(out, '\033', charBackspace)
			}
		}
	} else if isCtrl && !isAlt {
		// Ctrl+Key combinations
		if ker.wVirtualKeyCode == vkBack || ker.unicodeChar == vkBack {
			// Ctrl+Backspace -> CharCtrlW (kill word backward)
			out = append(out, charCtrlW)
		} else if ker.wVirtualKeyCode == vkDelete {
			// Ctrl+Delete -> MetaDelete (kill word forward)
			out = append(out, '\033', 'd')
		} else if ker.wVirtualKeyCode == vkLeft {
			// Ctrl+Left -> MetaBackward (jump word backward)
			out = append(out, '\033', 'b')
		} else if ker.wVirtualKeyCode == vkRight {
			// Ctrl+Right -> MetaForward (jump word forward)
			out = append(out, '\033', 'f')
		} else if ker.unicodeChar >= 1 && ker.unicodeChar <= 26 {
			out = append(out, byte(ker.unicodeChar))
		} else if ker.unicodeChar >= 'a' && ker.unicodeChar <= 'z' {
			out = append(out, byte(ker.unicodeChar-'a'+1))
		} else if ker.unicodeChar >= 'A' && ker.unicodeChar <= 'Z' {
			out = append(out, byte(ker.unicodeChar-'A'+1))
		} else if ker.unicodeChar == 0 && ker.wVirtualKeyCode >= 'A' && ker.wVirtualKeyCode <= 'Z' {
			out = append(out, byte(ker.wVirtualKeyCode-'A'+1))
		}
	} else {
		// Normal typing (!isAlt && !isCtrl, or AltGr)
		if ker.unicodeChar != 0 {
			ch := rune(ker.unicodeChar)
			switch ch {
			case '\r':
				out = append(out, '\r')
			case '\n':
				out = append(out, '\n')
			case '\b':
				out = append(out, '\b')
			case '\t':
				if isShift {
					out = append(out, '\033', '[', 'Z') // Shift+Tab (backtab)
				} else {
					out = append(out, '\t')
				}
			default:
				if utf16.IsSurrogate(ch) {
					if r.surrogate == 0 {
						r.surrogate = ch
						return
					}
					ch = utf16.DecodeRune(r.surrogate, ch)
					r.surrogate = 0
				} else {
					r.surrogate = 0
				}
				var b [4]byte
				n := utf8.EncodeRune(b[:], ch)
				out = append(out, b[:n]...)
			}
		} else {
			r.surrogate = 0
			switch ker.wVirtualKeyCode {
			case vkLeft:
				out = append(out, charBackward)
			case vkRight:
				out = append(out, charForward)
			case vkUp:
				out = append(out, charPrev)
			case vkDown:
				out = append(out, charNext)
			case vkHome:
				out = append(out, charLineStart)
			case vkEnd:
				out = append(out, charLineEnd)
			case vkDelete:
				out = append(out, '\033', '[', '3', '~')
			case vkBack:
				out = append(out, '\b')
			case vkTab:
				if isShift {
					out = append(out, '\033', '[', 'Z')
				} else {
					out = append(out, '\t')
				}
			case vkReturn:
				out = append(out, '\r')
			case vkEscape:
				out = append(out, charEsc)
			}
		}
	}

	if len(out) == 0 {
		return
	}

	repeat := int(ker.wRepeatCount)
	if repeat < 1 {
		repeat = 1
	}
	for i := 0; i < repeat; i++ {
		r.buf = append(r.buf, out...)
	}
}

func (r *consoleReader) Close() error {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil
	}
	r.closed = true
	r.mu.Unlock()

	// Post dummy focus event to unblock any pending ReadConsoleInputW
	var dummy inputRecord
	dummy.eventType = eventFocus
	var written uint32
	procWriteConsoleInputW.Call(
		uintptr(r.handle),
		uintptr(unsafe.Pointer(&dummy)),
		1,
		uintptr(unsafe.Pointer(&written)),
	)
	return nil
}
