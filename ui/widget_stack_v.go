package ui

type VStack struct {
	Children []Widget
}

func NewVStack(children ...Widget) VStack {
	return VStack{Children: children}
}

func (h VStack) Render(ctx Context, width, height int) {
	sizes := resolveConstraints(h.Children, height)
	start := 0
	for i, child := range h.Children {
		child.Render(ctx.SubContext(0, start), width, sizes[i])
		start += sizes[i]
	}
}
