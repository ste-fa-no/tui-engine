package ui

import (
	"tui-engine/backend"
	"tui-engine/renderer"
)

type UI struct {
	renderer *renderer.Renderer
}

func New(b backend.Backend) (*UI, error) {
	r, err := renderer.NewRenderer(b)
	if err != nil {
		return nil, err
	}
	return &UI{renderer: r}, nil
}

func (u *UI) Draw(widget Widget, x, y, width, height int) {
	ctx := newContext(u.renderer.CurrentBuffer(), x, y)
	widget.Render(ctx, width, height)
}

func (u *UI) Render() {
	u.renderer.Render()
}

func (u *UI) Resize(width, height int) {
	u.renderer.Resize(width, height)
}
