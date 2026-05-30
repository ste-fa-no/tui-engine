package ui

type ScrollableTextArea struct {
	TextArea
	Scrollbar
}

func (sta *ScrollableTextArea) Render(ctx Context, width, height int) {
	if width < 2 {
		sta.TextArea.Render(ctx, width, height)
		return
	}

	textWidth := width - 1

	layout := sta.TextArea.computeLayout(textWidth)
	sta.TextArea.clampOffset(layout, height)

	sta.Scrollbar.totalItems = len(layout.visualLines)
	sta.Scrollbar.visibleItems = min(height, len(layout.visualLines))
	sta.Scrollbar.offset = sta.TextArea.offsetRow

	sta.TextArea.Render(ctx.SubContext(0, 0), textWidth, height)
	sta.Scrollbar.Render(ctx.SubContext(textWidth, 0), 1, height)
}
