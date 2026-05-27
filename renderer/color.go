package renderer

import "fmt"

type Color struct {
	R, G, B uint8
}

var (
	ColorRed     = Color{255, 0, 0}
	ColorGreen   = Color{0, 255, 0}
	ColorBlue    = Color{0, 0, 255}
	ColorCyan    = Color{0, 255, 255}
	ColorMagenta = Color{255, 0, 255}
	ColorYellow  = Color{255, 255, 0}
	ColorBlack   = Color{0, 0, 0}
	ColorWhite   = Color{255, 255, 255}
)

func (c Color) SequenceForeground() string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c Color) SequenceBackground() string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
}
