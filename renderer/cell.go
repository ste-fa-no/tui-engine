package renderer

type Cell struct {
	Ch         rune
	Style      Style
	Foreground Color
	Background Color
}

type DrawnCell struct {
	X, Y int
	Cell Cell
}
