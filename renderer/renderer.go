package renderer

import (
	"fmt"
	"tui-engine/backend"
)

type Renderer struct {
	backend        backend.Backend
	previousBuffer *RenderBuffer
	currentBuffer  *RenderBuffer
}

func NewRenderer(backend backend.Backend) (*Renderer, error) {
	w, h, err := backend.Size()

	if err != nil {
		return nil, err
	}

	return &Renderer{
		backend:        backend,
		previousBuffer: NewRenderBuffer(w, h),
		currentBuffer:  NewRenderBuffer(w, h),
	}, nil
}

func (r *Renderer) Render() {
	for _, i := range r.previousBuffer.Diff(r.currentBuffer) {
		x := i % r.currentBuffer.Width
		y := i / r.currentBuffer.Width

		posSeq := fmt.Sprintf("\x1b[%d;%dH", y+1, x+1)

		cell := r.currentBuffer.Cells[i]
		data := []byte(posSeq +
			cell.Style.Sequence() +
			cell.Foreground.SequenceForeground() +
			cell.Background.SequenceBackground() +
			string(cell.Ch))

		err := r.backend.Write(data)
		if err != nil {
			return
		}
	}

	err := r.backend.Flush()
	if err != nil {
		return
	}
}

func (r *Renderer) Resize(width, height int) {
	r.previousBuffer = NewRenderBuffer(width, height)
	r.currentBuffer = NewRenderBuffer(width, height)

	err := r.backend.Write([]byte(backend.CLEAR_SCREEN))
	if err != nil {
		return
	}
	
	err = r.backend.Write([]byte(backend.CURSOR_ON_TOP))
	if err != nil {
		return
	}

	err = r.backend.Flush()
	if err != nil {
		return
	}
}

func (r *Renderer) CurrentBuffer() *RenderBuffer {
	return r.currentBuffer
}
