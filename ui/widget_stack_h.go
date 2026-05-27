package ui

type HStack struct {
	Children []Widget
}

func NewHStack(children ...Widget) HStack {
	return HStack{Children: children}
}

func (h HStack) Render(ctx Context, width, height int) {
	nc := len(h.Children)
	size := width / nc

	for i, child := range h.Children {
		child.Render(ctx.SubContext(size*i, 0), size, height)
	}
}
