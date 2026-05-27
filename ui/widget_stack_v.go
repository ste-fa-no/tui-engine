package ui

type VStack struct {
	Children []Widget
}

func NewVStack(children ...Widget) VStack {
	return VStack{Children: children}
}

func (h VStack) Render(ctx Context, width, height int) {
	nc := len(h.Children)
	size := height / nc

	for i, child := range h.Children {
		child.Render(ctx.SubContext(0, size*i), width, size)
	}
}
