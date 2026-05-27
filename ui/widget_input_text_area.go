package ui

import (
	"tui-engine/events"
	"tui-engine/renderer"
)

type TextArea struct {
	Input
	lines     []string
	cursorRow int
	cursorCol int
	offsetRow int
}

func (ta *TextArea) SetFocused(focused bool) {
	ta.focused = focused
}

func (ta *TextArea) HandleEvent(e events.Event) bool {
	if len(ta.lines) == 0 {
		ta.lines = []string{""}
	}

	switch ev := e.(type) {
	case events.KeyEvent:
		switch ev.Code {
		case events.KeyRune:
			line := []rune(ta.lines[ta.cursorRow])
			line = append(line[:ta.cursorCol], append([]rune{ev.Rune}, line[ta.cursorCol:]...)...)
			ta.lines[ta.cursorRow] = string(line)
			ta.cursorCol++

		case events.KeyEnter:
			line := ta.lines[ta.cursorRow]
			ta.lines[ta.cursorRow] = line[:ta.cursorCol]
			newLine := line[ta.cursorCol:]
			ta.lines = append(ta.lines[:ta.cursorRow+1],
				append([]string{newLine}, ta.lines[ta.cursorRow+1:]...)...)
			ta.cursorRow++
			ta.cursorCol = 0

		case events.KeyBackspace:
			if ta.cursorCol > 0 {
				line := []rune(ta.lines[ta.cursorRow])
				line = append(line[:ta.cursorCol-1], line[ta.cursorCol:]...)
				ta.lines[ta.cursorRow] = string(line)
				ta.cursorCol--
			} else if ta.cursorRow > 0 {
				prevLine := ta.lines[ta.cursorRow-1]
				currLine := ta.lines[ta.cursorRow]
				ta.cursorCol = len([]rune(prevLine))
				ta.lines[ta.cursorRow-1] = prevLine + currLine
				ta.lines = append(ta.lines[:ta.cursorRow], ta.lines[ta.cursorRow+1:]...)
				ta.cursorRow--
			}

		case events.KeyLeft:
			if ta.cursorCol > 0 {
				ta.cursorCol--
			} else if ta.cursorRow > 0 {
				ta.cursorRow--
				ta.cursorCol = len([]rune(ta.lines[ta.cursorRow]))
			}

		case events.KeyRight:
			if ta.cursorCol < len([]rune(ta.lines[ta.cursorRow])) {
				ta.cursorCol++
			} else if ta.cursorRow < len(ta.lines)-1 {
				ta.cursorRow++
				ta.cursorCol = 0
			}

		case events.KeyUp:
			if ta.cursorRow > 0 {
				ta.cursorRow--

				lineLen := len([]rune(ta.lines[ta.cursorRow]))
				if ta.cursorCol > lineLen {
					ta.cursorCol = lineLen
				}
			}

		case events.KeyDown:
			if ta.cursorRow < len(ta.lines)-1 {
				ta.cursorRow++

				lineLen := len([]rune(ta.lines[ta.cursorRow]))
				if ta.cursorCol > lineLen {
					ta.cursorCol = lineLen
				}
			}
		}

		return true
	}

	return false
}

type visualLine struct {
	text     string
	logRow   int
	colStart int
}

func (ta *TextArea) Render(ctx Context, width, height int) {
	if len(ta.lines) == 0 {
		ta.lines = []string{""}
	}

	// disegna placeholder se vuoto e non focalizzato
	if !ta.focused && len(ta.lines) == 1 && ta.lines[0] == "" {
		placeholderRunes := []rune(ta.Placeholder)
		for i := 0; i < width && i < len(placeholderRunes); i++ {
			ctx.Draw(i, 0, renderer.Cell{
				Ch:         placeholderRunes[i],
				Foreground: ta.PlaceholderForeground,
				Background: ta.Background,
			})
		}
		return
	}

	var visualLines []visualLine
	cursorVisualRow := 0
	cursorVisualCol := 0

	for logRow, line := range ta.lines {
		chunks := wrapLine(line, width)
		for chunkIdx, chunk := range chunks {
			colStart := chunkIdx * width
			vl := visualLine{
				text:     chunk,
				logRow:   logRow,
				colStart: colStart,
			}
			visualLines = append(visualLines, vl)

			if logRow == ta.cursorRow {
				if ta.cursorCol >= colStart && ta.cursorCol < colStart+width {
					cursorVisualRow = len(visualLines) - 1
					cursorVisualCol = ta.cursorCol - colStart
				} else if chunkIdx == len(chunks)-1 && ta.cursorCol >= colStart+len([]rune(chunk)) {
					cursorVisualRow = len(visualLines) - 1
					cursorVisualCol = len([]rune(chunk))
				}
			}
		}
	}

	if cursorVisualRow < ta.offsetRow {
		ta.offsetRow = cursorVisualRow
	}
	if cursorVisualRow >= ta.offsetRow+height {
		ta.offsetRow = cursorVisualRow - height + 1
	}
	if ta.offsetRow > 0 && ta.offsetRow+height > len(visualLines) {
		ta.offsetRow = len(visualLines) - height
		if ta.offsetRow < 0 {
			ta.offsetRow = 0
		}
	}

	for row := 0; row < height; row++ {
		visIdx := ta.offsetRow + row

		if visIdx >= len(visualLines) {
			for col := 0; col < width; col++ {
				ctx.Draw(col, row, renderer.Cell{
					Ch:         ' ',
					Foreground: ta.Foreground,
					Background: ta.Background,
				})
			}
			continue
		}

		runes := []rune(visualLines[visIdx].text)

		for col := 0; col < width; col++ {
			fg := ta.Foreground
			bg := ta.Background
			ch := ' '

			if col < len(runes) {
				ch = runes[col]
			}

			if ta.focused && visIdx == cursorVisualRow && col == cursorVisualCol {
				fg = ta.CursorForeground
				bg = ta.CursorBackground
			}

			ctx.Draw(col, row, renderer.Cell{
				Ch:         ch,
				Foreground: fg,
				Background: bg,
			})
		}
	}
}

func wrapLine(line string, width int) []string {
	runes := []rune(line)

	if len(runes) == 0 {
		return []string{""}
	}

	var result []string
	for len(runes) > 0 {
		if len(runes) <= width {
			result = append(result, string(runes))
			break
		}
		result = append(result, string(runes[:width]))
		runes = runes[width:]
	}

	return result
}
