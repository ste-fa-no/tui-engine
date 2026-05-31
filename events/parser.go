package events

var specialKeys = map[rune]KeyCode{
	'\r':   KeyEnter,
	'\x7f': KeyBackspace,
	'\x03': KeyCtrlC,
	'\x04': KeyCtrlD,
	'\t':   KeyTab,
	'\x1b': KeyEscape,
}

func parse(buf []byte) (Event, int) {
	char := rune(buf[0])
	keyCode, ok := specialKeys[char]

	if !ok {
		return KeyEvent{Code: KeyRune, Rune: char}, 1
	}

	if keyCode == KeyEscape {
		if len(buf) >= 3 && buf[1] == '[' {
			switch buf[2] {
			case 'A':
				keyCode = KeyUp
			case 'B':
				keyCode = KeyDown
			case 'C':
				keyCode = KeyRight
			case 'D':
				keyCode = KeyLeft
			case 'Z':
				keyCode = KeyShiftTab
			case '3':
				if len(buf) >= 4 && buf[3] == '~' {
					return KeyEvent{Code: KeyDelete}, 4
				}
			}
			return KeyEvent{Code: keyCode}, 3
		}

	}

	return KeyEvent{Code: keyCode}, 1
}
