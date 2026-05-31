package ui

type VStack struct {
	children []Widget
}

func NewVStack(children ...Widget) VStack {
	return VStack{children: children}
}

func (v VStack) Render(ctx Context, width, height int) {
	sizes := resolveConstraints(v.children, height)
	start := 0
	for i, child := range v.children {
		child.Render(ctx.SubContext(0, start), width, sizes[i])
		start += sizes[i]
	}
}

func (v VStack) Children() []Widget {
	return v.children
}
