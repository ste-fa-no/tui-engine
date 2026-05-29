package ui

import "tui-engine/renderer"

type Text struct {
	Content    string
	Foreground renderer.Color
	Background renderer.Color
	Style      renderer.Style
	Constraint Constraint
}

func (t Text) GetConstraint() Constraint {
	return t.Constraint
}

func (t Text) Render(ctx Context, width, height int) {
	col := 0
	row := 0
	for _, ch := range t.Content {
		if col >= width {
			col = 0
			row++
		}

		if row >= height {
			break
		}

		if ch == '\n' {
			for col < width {
				ctx.Draw(col, row, renderer.Cell{Ch: ' ', Foreground: t.Foreground, Background: t.Background})
				col++
			}
			col = 0
			row++
			continue
		}

		ctx.Draw(col, row, renderer.Cell{
			Ch:         ch,
			Style:      t.Style,
			Foreground: t.Foreground,
			Background: t.Background,
		})
		col++
	}

	for r := row; r < height; r++ {
		for c := col; c < width; c++ {
			ctx.Draw(c, r, renderer.Cell{Ch: ' ', Foreground: t.Foreground, Background: t.Background})
		}
		col = 0
	}
}
