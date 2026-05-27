package main

import (
	"tui-engine/backend"
	"tui-engine/events"
	"tui-engine/renderer"
	"tui-engine/ui"
)

func main() {
	b := backend.NewBackend()

	err := b.Init()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := b.Restore()
		if err != nil {
			panic(err)
		}
	}()

	u, err := ui.New(b)
	if err != nil {
		panic(err)
	}

	list := &ui.List{
		Items: []string{
			"Elemento 1",
			"Elemento 2",
			"Elemento 3",
			"Elemento 4",
			"Elemento 5",
			"Elemento 6",
			"Elemento 7",
			"Elemento 8",
			"Elemento 9",
			"Elemento 10",
			"Elemento 11",
			"Elemento 12",
			"Elemento 13",
			"Elemento 14",
			"Elemento 15",
		},
		Foreground:         renderer.ColorWhite,
		Background:         renderer.ColorBlack,
		SelectedForeground: renderer.ColorBlack,
		SelectedBackground: renderer.ColorGreen,
	}

	fm := ui.NewFocusManager()
	fm.Add(list)

	layout := ui.TitledBorder{
		Title:      "Lista con scrolling",
		Child:      list,
		Foreground: renderer.ColorWhite,
	}

	w, h, _ := b.Size()

	loop := events.NewEventLoop(b)
	ch := loop.Start()

	u.Draw(layout, 0, 0, w, h)
	u.Render()

	for e := range ch {
		switch ev := e.(type) {
		case events.KeyEvent:
			if ev.Code == events.KeyCtrlC {
				return
			}
			fm.HandleEvent(e)
		case events.ResizeEvent:
			w, h = ev.Width, ev.Height
			u.Resize(w, h)
		}

		u.Draw(layout, 0, 0, w, h)
		u.Render()
	}
}
