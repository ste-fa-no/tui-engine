//go:build windows

package backend

import (
	"golang.org/x/sys/windows"
)

type BackendWindows struct {
	oldStdoutMode uint32
	backendInit
	backendSize
	backendBuffer
	backendRestore
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

	err = b.backendInit.Init()

	return err
}

func (b *BackendWindows) Restore() error {
	err := b.backendRestore.Restore()

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
