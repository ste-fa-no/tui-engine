package ui

import (
	"tui-engine/actions"
)

type FocusManager struct {
	widgets []Interactive
	current int
}

func NewFocusManager() *FocusManager {
	return &FocusManager{}
}

func (fm *FocusManager) Add(w Interactive) {
	fm.widgets = append(fm.widgets, w)

	if len(fm.widgets) == 1 {
		w.SetFocused(true)
	}
}

func (fm *FocusManager) HandleAction(action actions.WidgetAction) bool {
	if len(fm.widgets) == 0 {
		return false
	}
	return fm.widgets[fm.current].HandleAction(action)
}

func (fm *FocusManager) HandleRune(r rune) bool {
	if len(fm.widgets) == 0 {
		return false
	}
	return fm.widgets[fm.current].HandleRune(r)
}

func (fm *FocusManager) FocusNext() {
	if len(fm.widgets) == 0 {
		return
	}
	fm.widgets[fm.current].SetFocused(false)
	fm.current = (fm.current + 1) % len(fm.widgets)
	fm.widgets[fm.current].SetFocused(true)
}

func (fm *FocusManager) FocusPrevious() {
	if len(fm.widgets) == 0 {
		return
	}
	fm.widgets[fm.current].SetFocused(false)
	fm.current = (fm.current - 1 + len(fm.widgets)) % len(fm.widgets)
	fm.widgets[fm.current].SetFocused(true)
}
