package app

import (
	"tui-engine/actions"
	"tui-engine/events"
	"tui-engine/ui"
)

type GlobalHandler func(action actions.GlobalAction)

type ActionDispatcher struct {
	keyMap        actions.KeyMap
	globalHandler GlobalHandler
	focusManager  *ui.FocusManager
}

func (ad *ActionDispatcher) Handle(e events.Event) {
	ev, ok := e.(events.KeyEvent)
	if !ok {
		return
	}

	combo := actions.KeyCombo{Code: ev.Code, Rune: ev.Rune}

	if action, ok := ad.keyMap.Global[combo]; ok {
		ad.globalHandler(action)
		return
	}

	if action, ok := ad.keyMap.Widget[combo]; ok {
		ad.focusManager.HandleAction(action)
		return
	}

	if ev.Code == events.KeyRune {
		ad.focusManager.HandleRune(ev.Rune)
	}
}
