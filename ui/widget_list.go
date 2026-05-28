package ui

import (
	"tui-engine/events"
	"tui-engine/renderer"
)

type List struct {
	Items              []string
	selected           int
	focused            bool
	offset             int
	Foreground         renderer.Color
	Background         renderer.Color
	SelectedForeground renderer.Color
	SelectedBackground renderer.Color
	Constraint         Constraint
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
	if l.selected < l.offset {
		l.offset = l.selected
	}
	if l.selected >= l.offset+height {
		l.offset = l.selected - height + 1
	}
	if l.offset > 0 && l.offset+height > len(l.Items) {
		l.offset = len(l.Items) - height
		if l.offset < 0 {
			l.offset = 0
		}
	}

	showScrollbar := len(l.Items) > height

	itemWidth := width
	if showScrollbar {
		itemWidth = width - 1
	}

	thumbHeight := 1
	thumbPos := 0
	if showScrollbar {
		thumbHeight = height * height / len(l.Items)
		if thumbHeight < 1 {
			thumbHeight = 1
		}
		thumbPos = l.offset * height / len(l.Items)

		if l.offset+height >= len(l.Items) {
			thumbPos = height - thumbHeight
		}
	}

	for i := 0; i < height; i++ {
		itemIdx := l.offset + i

		if itemIdx < len(l.Items) {
			s := l.Items[itemIdx]
			fg := l.Foreground
			bg := l.Background
			prefix := ' '

			if itemIdx == l.selected {
				fg = l.SelectedForeground
				bg = l.SelectedBackground
				prefix = '>'
			}

			ctx.Draw(0, i, renderer.Cell{Ch: prefix, Foreground: fg, Background: bg})

			for j, ch := range s {
				if j+2 >= itemWidth {
					break
				}
				ctx.Draw(j+2, i, renderer.Cell{Ch: ch, Foreground: fg, Background: bg})
			}

			lineLen := len([]rune(s)) + 2
			for j := lineLen; j < itemWidth; j++ {
				ctx.Draw(j, i, renderer.Cell{Ch: ' ', Foreground: fg, Background: bg})
			}
		} else {
			for j := 0; j < itemWidth; j++ {
				ctx.Draw(j, i, renderer.Cell{Ch: ' ', Foreground: l.Foreground, Background: l.Background})
			}
		}

		if showScrollbar {
			scrollCh := '│'
			if i >= thumbPos && i < thumbPos+thumbHeight {
				scrollCh = '█'
			}
			ctx.Draw(width-1, i, renderer.Cell{
				Ch:         scrollCh,
				Foreground: renderer.Color{R: 180, G: 180, B: 180},
				Background: l.Background,
			})
		}
	}
}

func (l *List) HandleEvent(e events.Event) bool {
	switch ev := e.(type) {
	case events.KeyEvent:
		if ev.Code == events.KeyUp {
			if l.selected > 0 {
				l.selected--
			}
		} else if ev.Code == events.KeyDown {
			if l.selected < len(l.Items)-1 {
				l.selected++
			}
		}
		return true
	}

	return false
}
