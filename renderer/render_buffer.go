package renderer

import "errors"

type RenderBuffer struct {
	Width, Height int
	Cells         []Cell
}

func NewRenderBuffer(width, height int) *RenderBuffer {
	return &RenderBuffer{Width: width, Height: height, Cells: make([]Cell, width*height)}
}

func (r *RenderBuffer) Get(x, y int) (Cell, error) {
	if x < 0 || y < 0 || x >= r.Width || y >= r.Height {
		return Cell{}, errors.New("out of bounds")
	}

	return r.Cells[y*r.Width+x], nil
}

func (r *RenderBuffer) Set(x, y int, cell Cell) error {
	if x < 0 || y < 0 || x >= r.Width || y >= r.Height {
		return errors.New("out of bounds")
	}

	r.Cells[y*r.Width+x] = cell
	return nil
}

func (r *RenderBuffer) Diff(other *RenderBuffer) []int {
	diffs := make([]int, 0, len(r.Cells))

	for i, cell := range r.Cells {
		if cell != other.Cells[i] {
			diffs = append(diffs, i)
		}
	}

	return diffs
}
