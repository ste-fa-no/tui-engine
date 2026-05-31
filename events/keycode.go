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
	KeyDelete
	KeyCtrlC
	KeyCtrlD
	KeyTab
	KeyShiftTab
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
	KeyDelete:    "DELETE",
	KeyCtrlC:     "CTRL-C",
	KeyCtrlD:     "CTRL-D",
	KeyTab:       "TAB",
	KeyShiftTab:  "SHIFT-TAB",
	KeyEscape:    "ESC",
}

func (code KeyCode) Name() string {
	return keyCodeNames[code]
}
