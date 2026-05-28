package events

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
