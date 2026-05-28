package events

import (
	"os"
)

type KeyEvent struct {
	Code KeyCode
	Rune rune
}

func (e KeyEvent) isEvent() {}

func watchKeyEvents(ch chan<- Event) {
	buffer := make([]byte, 16)
	for {
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			close(ch)
			return
		}
		event, _ := parse(buffer[:n])
		ch <- event
	}
}
