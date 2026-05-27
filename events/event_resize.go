package events

type ResizeEvent struct {
	Width, Height int
}

func (e ResizeEvent) isEvent() {}
