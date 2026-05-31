package ui

import (
	"tui-engine/actions"
	"tui-engine/renderer"
)

type TextArea struct {
	Input
	lines      []string
	cursorRow  int
	cursorCol  int
	offsetRow  int
	Constraint Constraint
}

func (ta *TextArea) IsFocused() bool {
	return ta.focused
}

func (ta *TextArea) SetFocused(focused bool) {
	ta.focused = focused
}

func (ta *TextArea) GetConstraint() Constraint {
	return ta.Constraint
}

func (ta *TextArea) HandleAction(action actions.WidgetAction) bool {
	if len(ta.lines) == 0 {
		ta.lines = []string{""}
	}

	switch action {
	case actions.CursorUp:
		if ta.cursorRow > 0 {
			ta.cursorRow--
			lineLen := len([]rune(ta.lines[ta.cursorRow]))
			if ta.cursorCol > lineLen {
				ta.cursorCol = lineLen
			}
		}

	case actions.CursorDown:
		if ta.cursorRow < len(ta.lines)-1 {
			ta.cursorRow++
			lineLen := len([]rune(ta.lines[ta.cursorRow]))
			if ta.cursorCol > lineLen {
				ta.cursorCol = lineLen
			}
		}

	case actions.CursorLeft:
		if ta.cursorCol > 0 {
			ta.cursorCol--
		} else if ta.cursorRow > 0 {
			ta.cursorRow--
			ta.cursorCol = len([]rune(ta.lines[ta.cursorRow]))
		}

	case actions.CursorRight:
		if ta.cursorCol < len([]rune(ta.lines[ta.cursorRow])) {
			ta.cursorCol++
		} else if ta.cursorRow < len(ta.lines)-1 {
			ta.cursorRow++
			ta.cursorCol = 0
		}

	case actions.Confirm:
		line := ta.lines[ta.cursorRow]
		ta.lines[ta.cursorRow] = line[:ta.cursorCol]
		newLine := line[ta.cursorCol:]
		ta.lines = append(ta.lines[:ta.cursorRow+1],
			append([]string{newLine}, ta.lines[ta.cursorRow+1:]...)...)
		ta.cursorRow++
		ta.cursorCol = 0

	case actions.DeleteBack:
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

	default:
		return false
	}

	return true
}

func (ta *TextArea) HandleRune(r rune) bool {
	if len(ta.lines) == 0 {
		ta.lines = []string{""}
	}

	line := []rune(ta.lines[ta.cursorRow])
	line = append(line[:ta.cursorCol], append([]rune{r}, line[ta.cursorCol:]...)...)
	ta.lines[ta.cursorRow] = string(line)
	ta.cursorCol++

	return true
}

func (ta *TextArea) Render(ctx Context, width, height int) {
	if len(ta.lines) == 0 {
		ta.lines = []string{""}
	}

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

	layout := ta.computeLayout(width)
	ta.clampOffset(layout, height)

	for row := 0; row < height; row++ {
		visIdx := ta.offsetRow + row

		if visIdx >= len(layout.visualLines) {
			for col := 0; col < width; col++ {
				ctx.Draw(col, row, renderer.Cell{
					Ch:         ' ',
					Foreground: ta.Foreground,
					Background: ta.Background,
				})
			}
			continue
		}

		runes := []rune(layout.visualLines[visIdx].text)

		for col := 0; col < width; col++ {
			fg := ta.Foreground
			bg := ta.Background
			ch := ' '

			if col < len(runes) {
				ch = runes[col]
			}

			if ta.focused && visIdx == layout.cursorVisualRow && col == layout.cursorVisualCol {
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

type visualLine struct {
	text     string
	logRow   int
	colStart int
}

type textAreaLayout struct {
	visualLines     []visualLine
	cursorVisualRow int
	cursorVisualCol int
}

func (ta *TextArea) computeLayout(width int) textAreaLayout {
	var visualLines []visualLine
	cursorVisualRow := 0
	cursorVisualCol := 0

	for logRow, line := range ta.lines {
		chunks := wrapLine(line, width)
		for chunkIdx, chunk := range chunks {
			colStart := chunkIdx * width
			visualLines = append(visualLines, visualLine{
				text:     chunk,
				logRow:   logRow,
				colStart: colStart,
			})

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

	return textAreaLayout{
		visualLines:     visualLines,
		cursorVisualRow: cursorVisualRow,
		cursorVisualCol: cursorVisualCol,
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

func (ta *TextArea) clampOffset(layout textAreaLayout, height int) {
	if layout.cursorVisualRow < ta.offsetRow {
		ta.offsetRow = layout.cursorVisualRow
	}
	if layout.cursorVisualRow >= ta.offsetRow+height {
		ta.offsetRow = layout.cursorVisualRow - height + 1
	}
	if ta.offsetRow > 0 && ta.offsetRow+height > len(layout.visualLines) {
		ta.offsetRow = len(layout.visualLines) - height
		if ta.offsetRow < 0 {
			ta.offsetRow = 0
		}
	}
}
