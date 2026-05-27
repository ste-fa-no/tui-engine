package renderer

type Canvas interface {
	Set(x, y int, cell Cell) error
}
