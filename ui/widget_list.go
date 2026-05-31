package ui

import (
	"tui-engine/actions"
	"tui-engine/renderer"
	"tui-engine/util"
)

type List struct {
	Items              []string
	cursor             int
	focused            bool
	offset             int
	Foreground         renderer.Color
	Background         renderer.Color
	SelectedForeground renderer.Color
	SelectedBackground renderer.Color
	Constraint         Constraint
	Selection          util.State[int]
}

func (l *List) IsFocused() bool {
	return l.focused
}

func (l *List) SetFocused(focused bool) {
	l.focused = focused
}

func (l *List) GetConstraint() Constraint {
	return l.Constraint
}

func (l *List) Render(ctx Context, width, height int) {
	l.clampOffset(height)

	for i := 0; i < height && l.offset+i < len(l.Items); i++ {
		itemIdx := l.offset + i
		s := l.Items[itemIdx]

		fg := l.Foreground
		bg := l.Background
		prefix := ' '

		if itemIdx == l.cursor {
			fg = l.SelectedForeground
			bg = l.SelectedBackground
			prefix = '>'
		}

		ctx.Draw(0, i, renderer.Cell{Ch: prefix, Foreground: fg, Background: bg})

		for j, ch := range s {
			if j+2 >= width {
				break
			}
			ctx.Draw(j+2, i, renderer.Cell{Ch: ch, Foreground: fg, Background: bg})
		}

		lineLen := len([]rune(s)) + 2
		for j := lineLen; j < width; j++ {
			ctx.Draw(j, i, renderer.Cell{Ch: ' ', Foreground: fg, Background: bg})
		}
	}

	visibleItems := len(l.Items) - l.offset
	if visibleItems > height {
		visibleItems = height
	}
	for i := visibleItems; i < height; i++ {
		for j := 0; j < width; j++ {
			ctx.Draw(j, i, renderer.Cell{Ch: ' ', Foreground: l.Foreground, Background: l.Background})
		}
	}
}

func (l *List) HandleAction(action actions.WidgetAction) bool {
	switch action {
	case actions.CursorUp:
		if l.cursor > 0 {
			l.cursor--
		}
	case actions.CursorDown:
		if l.cursor < len(l.Items)-1 {
			l.cursor++
		}
	case actions.Confirm:
		l.Selection.Set(l.cursor)
	default:
		return false
	}
	return false
}

func (l *List) HandleRune(r rune) bool {
	return false
}

func (l *List) Offset() int {
	return l.offset
}

func (l *List) clampOffset(height int) {
	if l.cursor < l.offset {
		l.offset = l.cursor
	}
	if l.cursor >= l.offset+height {
		l.offset = l.cursor - height + 1
	}
	if l.offset > 0 && l.offset+height > len(l.Items) {
		l.offset = len(l.Items) - height
		if l.offset < 0 {
			l.offset = 0
		}
	}
}
