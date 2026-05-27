//go:build windows

package backend

import (
	"os"

	"golang.org/x/sys/windows"
	"golang.org/x/term"
)

type BackendWindows struct {
	oldState      *term.State
	oldStdoutMode uint32
	backendSize
	backendBuffer
}

func NewBackend() Backend {
	return &BackendWindows{}
}

func (b *BackendWindows) Init() error {
	handle := windows.Handle(STDOUT)

	var mode uint32
	err := windows.GetConsoleMode(handle, &mode)
	if err != nil {
		return err
	}

	b.oldStdoutMode = mode
	err = windows.SetConsoleMode(handle, mode|0x0004)

	if err != nil {
		return err
	}

	state, err := term.MakeRaw(int(STDIN))

	if err != nil {
		return err
	}

	b.oldState = state
	_, err = os.Stdout.WriteString(ALTERNATE_SCREEN)

	if err != nil {
		return err
	}

	_, err = os.Stdout.WriteString(HIDE_CURSOR)

	if err != nil {
		return err
	}

	return nil
}

func (b *BackendWindows) Restore() error {
	_, err := os.Stdout.WriteString(RESTORE_SCREEN)

	if err != nil {
		return err
	}

	_, err = os.Stdout.WriteString(SHOW_CURSOR)

	if err != nil {
		return err
	}

	err = term.Restore(int(STDIN), b.oldState)

	if err != nil {
		return err
	}

	handle := windows.Handle(STDOUT)
	err = windows.SetConsoleMode(handle, b.oldStdoutMode)

	if err != nil {
		return err
	}

	return nil
}
