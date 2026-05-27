//go:build !windows

package backend

import (
	"os"

	"golang.org/x/term"
)

type BackendUnix struct {
	oldState *term.State
	backendSize
	backendBuffer
}

func NewBackend() Backend {
	return &BackendUnix{}
}

func (b *BackendUnix) Init() error {
	state, err := term.MakeRaw(int(STDIN))

	if err != nil {
		return err
	}

	b.oldState = state
	os.Stdout.WriteString(ALTERNATE_SCREEN)
	os.Stdout.WriteString(HIDE_CURSOR)
	return err
}

func (b *BackendUnix) Restore() error {
	os.Stdout.WriteString(RESTORE_SCREEN)
	os.Stdout.WriteString(SHOW_CURSOR)

	err := term.Restore(int(STDIN), b.oldState)
	return err
}
