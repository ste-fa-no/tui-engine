package events

import "os"

type KeyCode int

const (
	KeyRune KeyCode = iota
	KeyUp
	KeyDown
	KeyLeft
	KeyRight
	KeyEnter
	KeyBackspace
	KeyCtrlC
	KeyCtrlD
	KeyTab
	KeyEscape
)

var keyCodeNames = map[KeyCode]string{
	KeyRune:      "RUNE",
	KeyUp:        "UP",
	KeyDown:      "DOWN",
	KeyLeft:      "LEFT",
	KeyRight:     "RIGHT",
	KeyEnter:     "ENTER",
	KeyBackspace: "BACKSPACE",
	KeyCtrlC:     "CTRL-C",
	KeyCtrlD:     "CTRL-D",
	KeyTab:       "TAB",
	KeyEscape:    "ESC",
}

func (code KeyCode) Name() string {
	return keyCodeNames[code]
}

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
