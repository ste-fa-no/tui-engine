package ui

import "tui-engine/renderer"

type Context interface {
	Draw(x, y int, cell renderer.Cell)
	SubContext(offsetX, offsetY int) Context
}

type context struct {
	canvas           renderer.Canvas
	offsetX, offsetY int
}

func newContext(canvas renderer.Canvas, offsetX, offsetY int) *context {
	return &context{
		canvas:  canvas,
		offsetX: offsetX,
		offsetY: offsetY,
	}
}

func (c *context) Draw(x, y int, cell renderer.Cell) {
	err := c.canvas.Set(c.offsetX+x, c.offsetY+y, cell)
	if err != nil {
		panic(err)
	}
}

func (c *context) SubContext(offsetX, offsetY int) Context {
	return newContext(
		c.canvas,
		c.offsetX+offsetX,
		c.offsetY+offsetY,
	)
}
