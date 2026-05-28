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
	AlternateScreen = "\x1b[?1049h"
	RestoreScreen   = "\x1b[?1049l"

	HideCursor = "\x1b[?25l"
	ShowCursor = "\x1b[?25h"

	ClearScreen = "\x1b[2J"
	CursorOnTop = "\x1b[H"
)

type Backend interface {
	Init() error
	Restore() error
	Size() (width int, height int, err error)
	Write(data []byte) error
	Flush() error
}

type baseBackend struct {
	buffer   bytes.Buffer
	oldState *term.State
}

func (b *baseBackend) Init() error {
	state, err := term.MakeRaw(int(STDIN))

	if err != nil {
		return err
	}

	b.oldState = state

	_, err = os.Stdout.WriteString(AlternateScreen)

	if err != nil {
		return err
	}

	_, err = os.Stdout.WriteString(HideCursor)

	return err
}

func (b *baseBackend) Restore() error {
	_, err := os.Stdout.WriteString(RestoreScreen)

	if err != nil {
		return err
	}

	_, err = os.Stdout.WriteString(ShowCursor)

	if err != nil {
		return err
	}

	err = term.Restore(int(STDIN), b.oldState)

	return err
}

func (b *baseBackend) Size() (width int, height int, err error) {
	return term.GetSize(int(STDOUT))
}

func (b *baseBackend) Write(data []byte) error {
	b.buffer.Write(data)
	return nil
}

func (b *baseBackend) Flush() error {
	_, err := b.buffer.WriteTo(os.Stdout)
	return err
}
