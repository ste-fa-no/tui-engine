package events

import (
	"os"
	"time"

	"golang.org/x/term"
)

func watchResizeEvents(ch chan<- Event) {
	prevW, prevH, _ := term.GetSize(int(os.Stdout.Fd()))

	for {
		time.Sleep(200 * time.Millisecond)

		w, h, err := term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			return
		}

		if w != prevW || h != prevH {
			ch <- ResizeEvent{Width: w, Height: h}
			prevW, prevH = w, h
		}
	}

}
