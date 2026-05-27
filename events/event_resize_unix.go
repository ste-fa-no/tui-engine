//go:build !windows

package events

import (
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/term"
)

func watchResizeEvents(ch chan<- Event) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGWINCH)

	for range sigCh {
		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			return
		}
		ch <- ResizeEvent{Width: w, Height: h}
	}
}
