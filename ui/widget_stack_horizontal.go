package ui

type HStack struct {
	Children []Widget
}

func NewHStack(children ...Widget) HStack {
	return HStack{Children: children}
}

func (h HStack) Render(ctx Context, width, height int) {
	sizes := resolveConstraints(h.Children, width)
	start := 0
	for i, child := range h.Children {
		child.Render(ctx.SubContext(start, 0), sizes[i], height)
		start += sizes[i]
	}
}
