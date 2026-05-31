package main

import (
	"tui-engine/app"
	"tui-engine/renderer"
	"tui-engine/ui"
)

func main() {
	list := &ui.ScrollableList{
		List: ui.List{
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
		},
		Scrollbar: ui.DefaultScrollbar,
	}

	textarea := &ui.ScrollableTextArea{
		TextArea: ui.TextArea{
			Input: ui.Input{
				Foreground:            renderer.ColorWhite,
				Background:            renderer.ColorBlack,
				CursorForeground:      renderer.ColorBlack,
				CursorBackground:      renderer.ColorWhite,
				PlaceholderForeground: renderer.Color{R: 100, G: 100, B: 100},
				Placeholder:           "Scrivi qualcosa...",
			},
		},
		Scrollbar: ui.DefaultScrollbar,
	}

	listBorder := &ui.TitledBorder{
		Title:             "ScrollableList",
		Child:             list,
		Foreground:        renderer.ColorWhite,
		ForegroundFocused: renderer.ColorGreen,
	}

	textareaBorder := &ui.TitledBorder{
		Title:             "ScrollableTextArea",
		Child:             textarea,
		Foreground:        renderer.ColorWhite,
		ForegroundFocused: renderer.ColorGreen,
	}

	layout := ui.NewHStack(listBorder, textareaBorder)

	a, err := app.New(&layout)
	if err != nil {
		panic(err)
	}

	if err := a.Run(); err != nil {
		panic(err)
	}
}
