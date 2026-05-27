package ui

import "tui-engine/events"

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

func (fm *FocusManager) HandleEvent(e events.Event) {
	if len(fm.widgets) == 0 {
		return
	}
	
	switch ev := e.(type) {
	case events.KeyEvent:
		if ev.Code == events.KeyTab {
			fm.widgets[fm.current].SetFocused(false)
			fm.current = (fm.current + 1) % len(fm.widgets)
			fm.widgets[fm.current].SetFocused(true)
		} else {
			fm.widgets[fm.current].HandleEvent(ev)
		}
	}
}
