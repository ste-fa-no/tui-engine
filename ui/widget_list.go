package ui

import (
	"tui-engine/events"
	"tui-engine/renderer"
)

type List struct {
	Items    []string
	selected int
	focused  bool
	offset   int

	Foreground renderer.Color
	Background renderer.Color

	SelectedForeground renderer.Color
	SelectedBackground renderer.Color
}

func (t *List) SetFocused(focused bool) {
	t.focused = focused
}

func (t *List) Render(ctx Context, width, height int) {
	if t.selected < t.offset {
		t.offset = t.selected
	}
	if t.selected >= t.offset+height {
		t.offset = t.selected - height + 1
	}
	if t.offset > 0 && t.offset+height > len(t.Items) {
		t.offset = len(t.Items) - height
		if t.offset < 0 {
			t.offset = 0
		}
	}

	showScrollbar := len(t.Items) > height

	itemWidth := width
	if showScrollbar {
		itemWidth = width - 1
	}

	thumbHeight := 1
	thumbPos := 0
	if showScrollbar {
		thumbHeight = height * height / len(t.Items)
		if thumbHeight < 1 {
			thumbHeight = 1
		}
		thumbPos = t.offset * height / len(t.Items)
	}

	for i := 0; i < height; i++ {
		itemIdx := t.offset + i

		if itemIdx < len(t.Items) {
			s := t.Items[itemIdx]
			fg := t.Foreground
			bg := t.Background
			prefix := ' '

			if itemIdx == t.selected {
				fg = t.SelectedForeground
				bg = t.SelectedBackground
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
				ctx.Draw(j, i, renderer.Cell{Ch: ' ', Foreground: t.Foreground, Background: t.Background})
			}
		}

		if showScrollbar {
			scrollCh := '│'
			if i >= thumbPos && i < thumbPos+thumbHeight {
				scrollCh = '█'
			}
			ctx.Draw(width-1, i, renderer.Cell{
				Ch:         rune(scrollCh),
				Foreground: renderer.Color{R: 180, G: 180, B: 180},
				Background: t.Background,
			})
		}
	}
}

func (t *List) HandleEvent(e events.Event) bool {
	switch ev := e.(type) {
	case events.KeyEvent:
		if ev.Code == events.KeyUp {
			if t.selected > 0 {
				t.selected--
			}
		} else if ev.Code == events.KeyDown {
			if t.selected < len(t.Items)-1 {
				t.selected++
			}
		}
		return true
	}

	return false
}
