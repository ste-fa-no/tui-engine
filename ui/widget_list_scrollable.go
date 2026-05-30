package ui

type ScrollableList struct {
	List
	Scrollbar
}

func (sl *ScrollableList) Render(ctx Context, width, height int) {
	if width < 2 {
		sl.List.Render(ctx, width, height)
		return
	}

	sl.List.clampOffset(height)

	sl.Scrollbar.totalItems = len(sl.List.Items)
	sl.Scrollbar.visibleItems = min(height, len(sl.List.Items))
	sl.Scrollbar.offset = sl.List.Offset()

	listWidth := width - 1
	sl.List.Render(ctx.SubContext(0, 0), listWidth, height)
	sl.Scrollbar.Render(ctx.SubContext(listWidth, 0), 1, height)
}
