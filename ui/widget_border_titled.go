package ui

import "tui-engine/renderer"

type TitledBorder struct {
	Child      Widget
	Foreground renderer.Color
	Background renderer.Color
	Title      string
}

func (b TitledBorder) Render(ctx Context, width, height int) {
	titleRunes := []rune(b.Title)
	var ch rune
	for i := 0; i < width; i++ {
		titleIdx := i - 2
		for j := 0; j < height; j++ {
			if i == j && i == 0 {
				ch = '┌'
			} else if i == width-1 && j == 0 {
				ch = '┐'
			} else if i == 0 && j == height-1 {
				ch = '└'
			} else if i == width-1 && j == height-1 {
				ch = '┘'
			} else if i == 0 || i == width-1 {
				ch = '│'
			} else if j == height-1 {
				ch = '─'
			} else if j == 0 {
				if b.Title != "" && titleIdx >= 0 && titleIdx < len(titleRunes) {
					ch = titleRunes[titleIdx]
				} else {
					ch = '─'
				}
			} else {
				continue
			}

			ctx.Draw(
				i,
				j,
				renderer.Cell{
					Ch:         ch,
					Foreground: b.Foreground,
					Background: b.Background},
			)
		}
	}

	b.Child.Render(ctx.SubContext(1, 1), width-2, height-2)
}
