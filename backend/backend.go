package backend

import (
	"bytes"
	"os"

	"golang.org/x/term"
)

var (
	STDIN  = os.Stdin.Fd()
	STDOUT = os.Stdout.Fd()
)

const (
	ALTERNATE_SCREEN = "\x1b[?1049h"
	RESTORE_SCREEN   = "\x1b[?1049l"

	HIDE_CURSOR = "\x1b[?25l"
	SHOW_CURSOR = "\x1b[?25h"

	CLEAR_SCREEN  = "\x1b[2J"
	CURSOR_ON_TOP = "\x1b[H"
)

type Backend interface {
	Init() error
	Restore() error
	Size() (width int, height int, err error)
	Write(data []byte) error
	Flush() error
}

type backendSize struct {
}

func (b backendSize) Size() (width int, height int, err error) {
	return term.GetSize(int(STDOUT))
}

type backendBuffer struct {
	buffer bytes.Buffer
}

func (b *backendBuffer) Write(data []byte) error {
	b.buffer.Write(data)
	return nil
}

func (b *backendBuffer) Flush() error {
	_, err := b.buffer.WriteTo(os.Stdout)
	return err
}
