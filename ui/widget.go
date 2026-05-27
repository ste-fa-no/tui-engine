package ui

import (
	"tui-engine/events"
	"tui-engine/renderer"
)

type Context interface {
	Draw(x, y int, cell renderer.Cell)
	SubContext(offsetX, offsetY int) Context
}

type Widget interface {
	Render(ctx Context, width, height int)
}

type Interactive interface {
	Widget
	SetFocused(focused bool)
	HandleEvent(event events.Event) bool
}

type Input struct {
	Value                 string
	Placeholder           string
	Foreground            renderer.Color
	Background            renderer.Color
	CursorForeground      renderer.Color
	CursorBackground      renderer.Color
	PlaceholderForeground renderer.Color
	focused               bool
	cursor                int
	offset                int
}
