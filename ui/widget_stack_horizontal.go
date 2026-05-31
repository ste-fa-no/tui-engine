package ui

type HStack struct {
	children []Widget
}

func NewHStack(children ...Widget) HStack {
	return HStack{children: children}
}

func (h HStack) Render(ctx Context, width, height int) {
	sizes := resolveConstraints(h.children, width)
	start := 0
	for i, child := range h.children {
		child.Render(ctx.SubContext(start, 0), sizes[i], height)
		start += sizes[i]
	}
}

func (h HStack) Children() []Widget {
	return h.children
}
