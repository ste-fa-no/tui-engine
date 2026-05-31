package ui

import (
	"tui-engine/actions"
	_ "tui-engine/events"
	"tui-engine/renderer"
)

type Widget interface {
	Render(ctx Context, width, height int)
}

type Interactive interface {
	Widget
	IsFocused() bool
	SetFocused(focused bool)
	HandleAction(action actions.WidgetAction) bool
	HandleRune(r rune) bool
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

type Container interface {
	Widget
	Children() []Widget
}
