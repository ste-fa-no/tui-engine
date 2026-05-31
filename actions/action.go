package actions

type Action interface {
	isAction()
}

type GlobalAction int
type WidgetAction int

const (
	AppQuit GlobalAction = iota
	FocusPrevious
	FocusNext
)

const (
	CursorUp WidgetAction = iota
	CursorDown
	CursorLeft
	CursorRight
	CursorLineStart
	CursorLineEnd
	Confirm
	Cancel
	DeleteBack
	DeleteForward
	NewLine
)

func (g GlobalAction) isAction() {}
func (w WidgetAction) isAction() {}
