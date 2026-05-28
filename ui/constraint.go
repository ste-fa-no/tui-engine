package ui

type Constraint interface {
	isConstraint()
}

type Fill struct {
}

func (f Fill) isConstraint() {}

type Fixed struct {
	Size int
}

func (f Fixed) isConstraint() {}

type Percent struct {
	Value int
}

func (p Percent) isConstraint() {}

type Constrainable interface {
	GetConstraint() Constraint
}

func resolveConstraints(children []Widget, available int) []int {
	sizes := make([]int, len(children))
	remaining := available
	fillCount := 0

	for i, child := range children {
		constraint := Constraint(Fill{})
		if c, ok := child.(Constrainable); ok {
			if c.GetConstraint() != nil {
				constraint = c.GetConstraint()
			}
		}

		switch c := constraint.(type) {
		case Fixed:
			sizes[i] = c.Size
			remaining -= c.Size
		case Percent:
			size := available * c.Value / 100
			sizes[i] = size
			remaining -= size
		default:
			fillCount++
		}
	}

	if fillCount > 0 {
		fillSize := remaining / fillCount
		for i, child := range children {
			if _, ok := child.(Constrainable); !ok {
				sizes[i] = fillSize
				continue
			}
			c := child.(Constrainable).GetConstraint()
			if _, ok := c.(Fixed); ok {
				continue
			}
			if _, ok := c.(Percent); ok {
				continue
			}
			sizes[i] = fillSize
		}
	}

	return sizes
}
