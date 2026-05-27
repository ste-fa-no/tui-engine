package ui

import "tui-engine/renderer"

type context struct {
	canvas           renderer.Canvas
	offsetX, offsetY int
}

func newContext(canvas renderer.Canvas, offsetX, offsetY int) *context {
	return &context{canvas: canvas, offsetX: offsetX, offsetY: offsetY}
}

func (c *context) Draw(x, y int, cell renderer.Cell) {
	err := c.canvas.Set(c.offsetX+x, c.offsetY+y, cell)
	if err != nil {
		panic(err)
	}
}

func (c *context) SubContext(offsetX, offsetY int) Context {
	return &context{
		canvas:  c.canvas,
		offsetX: c.offsetX + offsetX,
		offsetY: c.offsetY + offsetY,
	}
}
