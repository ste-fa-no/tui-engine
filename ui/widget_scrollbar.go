package ui

import "tui-engine/renderer"

type ScrollbarStyle struct {
	Track rune
	Thumb rune

	TrackForeground renderer.Color
	ThumbForeground renderer.Color

	TrackBackground renderer.Color
	ThumbBackground renderer.Color
}

type Scrollbar struct {
	Style ScrollbarStyle

	totalItems   int
	visibleItems int
	offset       int
}

var DefaultScrollbarStyle = ScrollbarStyle{
	Track: '│',
	Thumb: '█',

	TrackForeground: renderer.Color{R: 80, G: 80, B: 80},
	ThumbForeground: renderer.Color{R: 180, G: 180, B: 180},

	TrackBackground: renderer.Color{},
	ThumbBackground: renderer.Color{},
}

var DefaultScrollbar = Scrollbar{
	Style: DefaultScrollbarStyle,
}

func (sb *Scrollbar) Render(ctx Context, width, height int) {
	if sb.totalItems <= sb.visibleItems {
		for i := range height {
			ctx.Draw(0, i, renderer.Cell{
				Ch:         sb.Style.Track,
				Foreground: sb.Style.TrackForeground,
				Background: sb.Style.TrackBackground,
			})
		}
		return
	}

	thumbSize := max(1, height*sb.visibleItems/sb.totalItems)
	maxOffset := sb.totalItems - sb.visibleItems
	thumbPos := sb.offset * (height - thumbSize) / maxOffset
	if sb.offset >= maxOffset {
		thumbPos = height - thumbSize
	}

	for i := range height {
		cell := renderer.Cell{
			Ch:         sb.Style.Track,
			Foreground: sb.Style.TrackForeground,
			Background: sb.Style.TrackBackground,
		}
		if i >= thumbPos && i < thumbPos+thumbSize {
			cell.Ch = sb.Style.Thumb
			cell.Foreground = sb.Style.ThumbForeground
			cell.Background = sb.Style.ThumbBackground
		}
		ctx.Draw(0, i, cell)
	}
}
