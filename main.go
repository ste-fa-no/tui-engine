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
		Items:              []string{"Elemento 1", "Elemento 2", "Elemento 3", "Elemento 4", "Elemento 5"},
		Foreground:         renderer.ColorWhite,
		Background:         renderer.ColorBlack,
		SelectedForeground: renderer.ColorBlack,
		SelectedBackground: renderer.ColorGreen,
	}

	textArea := &ui.TextArea{
		Input: ui.Input{
			Placeholder:           "Scrivi qui...",
			Foreground:            renderer.ColorWhite,
			Background:            renderer.ColorBlack,
			CursorForeground:      renderer.ColorBlack,
			CursorBackground:      renderer.ColorWhite,
			PlaceholderForeground: renderer.Color{R: 128, G: 128, B: 128},
		},
	}

	panel := ui.Text{
		Content:    "Pannello al 30%",
		Foreground: renderer.ColorYellow,
		Background: renderer.ColorBlack,
	}

	fm := ui.NewFocusManager()
	fm.Add(list)
	fm.Add(textArea)

	layout := ui.NewHStack(
		ui.TitledBorder{
			Title:             "Lista (Fixed 30)",
			Child:             list,
			Foreground:        renderer.ColorWhite,
			ForegroundFocused: renderer.ColorGreen,
			Constraint:        ui.Fixed{Size: 30},
		},
		ui.TitledBorder{
			Title:             "Note (Fill)",
			Child:             textArea,
			Foreground:        renderer.ColorWhite,
			ForegroundFocused: renderer.ColorGreen,
			Constraint:        ui.Fill{},
		},
		ui.TitledBorder{
			Title:      "Panel (30%)",
			Child:      panel,
			Foreground: renderer.ColorWhite,
			Constraint: ui.Percent{Value: 30},
		},
	)

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
