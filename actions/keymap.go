package actions

import "tui-engine/events"

type KeyCombo struct {
	Code events.KeyCode
	Rune rune
}

type KeyMap struct {
	Global map[KeyCombo]GlobalAction
	Widget map[KeyCombo]WidgetAction
}

var DefaultKeyMap = KeyMap{
	Global: map[KeyCombo]GlobalAction{
		{Code: events.KeyCtrlC}:    AppQuit,
		{Code: events.KeyTab}:      FocusNext,
		{Code: events.KeyShiftTab}: FocusPrevious,
	},
	Widget: map[KeyCombo]WidgetAction{
		{Code: events.KeyUp}:        CursorUp,
		{Code: events.KeyDown}:      CursorDown,
		{Code: events.KeyLeft}:      CursorLeft,
		{Code: events.KeyRight}:     CursorRight,
		{Code: events.KeyEnter}:     Confirm,
		{Code: events.KeyEscape}:    Cancel,
		{Code: events.KeyBackspace}: DeleteBack,
		{Code: events.KeyDelete}:    DeleteForward,
	},
}
