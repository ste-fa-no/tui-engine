package ui

import "tui-engine/renderer"

type Border struct {
	Child      Widget
	Foreground renderer.Color
	Background renderer.Color
	Constraint Constraint
}

func (b Border) GetConstraint() Constraint {
	return b.Constraint
}

func (b Border) Render(ctx Context, width, height int) {
	var ch rune
	for i := 0; i < width; i++ {
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
			} else if j == 0 || j == height-1 {
				ch = '─'
			} else {
				continue
			}

			ctx.Draw(
				i, j,
				renderer.Cell{
					Ch:         ch,
					Foreground: b.Foreground,
					Background: b.Background},
			)

		}
	}

	b.Child.Render(ctx.SubContext(1, 1), width-2, height-2)
}
