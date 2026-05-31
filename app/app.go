package app

import (
	"tui-engine/actions"
	"tui-engine/backend"
	"tui-engine/events"
	"tui-engine/ui"
)

type App struct {
	backend      backend.Backend
	ui           *ui.UI
	dispatcher   *ActionDispatcher
	focusManager *ui.FocusManager
	root         ui.Widget
	quit         chan struct{}
}

func New(root ui.Widget) (*App, error) {
	b := backend.NewBackend()

	u, err := ui.New(b)
	if err != nil {
		return nil, err
	}

	fm := ui.NewFocusManager()

	a := &App{
		backend:      b,
		ui:           u,
		focusManager: fm,
		root:         root,
		quit:         make(chan struct{}),
	}

	gh := func(action actions.GlobalAction) {
		switch action {
		case actions.AppQuit:
			close(a.quit)
		case actions.FocusNext:
			fm.FocusNext()
		case actions.FocusPrevious:
			fm.FocusPrevious()
		}
	}

	a.dispatcher = &ActionDispatcher{
		actions.DefaultKeyMap,
		gh,
		fm,
	}

	a.collectInteractive(root)

	return a, nil
}

func (a *App) Run() error {
	err := a.backend.Init()
	if err != nil {
		return err
	}

	defer func() {
		_ = a.backend.Restore()
	}()

	w, h, _ := a.backend.Size()

	loop := events.NewEventLoop(a.backend)
	ch := loop.Start()

	a.ui.Draw(a.root, 0, 0, w, h)
	a.ui.Render()

	for {
		select {
		case <-a.quit:
			return nil
		case e := <-ch:
			switch ev := e.(type) {
			case events.KeyEvent:
				a.dispatcher.Handle(ev)
			case events.ResizeEvent:
				w, h = ev.Width, ev.Height
				a.ui.Resize(w, h)
			}

			a.ui.Draw(a.root, 0, 0, w, h)
			a.ui.Render()
		}
	}
}

func (a *App) collectInteractive(w ui.Widget) {
	if i, ok := w.(ui.Interactive); ok {
		a.focusManager.Add(i)
	}
	if c, ok := w.(ui.Container); ok {
		for _, child := range c.Children() {
			a.collectInteractive(child)
		}
	}
}
