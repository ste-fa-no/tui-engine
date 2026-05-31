package ui

import (
	"tui-engine/actions"
	"tui-engine/renderer"
)

type TextField struct {
	Input
	Constraint Constraint
}

func (tf *TextField) IsFocused() bool {
	return tf.focused
}

func (tf *TextField) SetFocused(focused bool) {
	tf.focused = focused
}

func (tf *TextField) GetConstraint() Constraint {
	return tf.Constraint
}

func (tf *TextField) HandleAction(action actions.WidgetAction) bool {
	switch action {
	case actions.CursorLeft:
		if tf.cursor > 0 {
			tf.cursor--
		}
	case actions.CursorRight:
		if tf.cursor < len([]rune(tf.Value)) {
			tf.cursor++
		}
	case actions.DeleteBack:
		if tf.cursor > 0 {
			runes := []rune(tf.Value)
			runes = append(runes[:tf.cursor-1], runes[tf.cursor:]...)
			tf.Value = string(runes)
			tf.cursor--
		}
	default:
		return false
	}
	return true
}

func (tf *TextField) HandleRune(r rune) bool {
	runes := []rune(tf.Value)
	runes = append(runes[:tf.cursor], append([]rune{r}, runes[tf.cursor:]...)...)
	tf.Value = string(runes)
	tf.cursor++
	return true
}

func (tf *TextField) Render(ctx Context, width, height int) {
	runes := []rune(tf.Value)

	if tf.cursor > len(runes) {
		tf.cursor = len(runes)
	}

	if tf.cursor < tf.offset {
		tf.offset = tf.cursor
	}

	if tf.cursor >= tf.offset+width {
		tf.offset = tf.cursor - width + 1
	}

	if tf.offset > 0 && tf.offset+width > len(runes)+1 {
		tf.offset = len(runes) - width + 1
		if tf.offset < 0 {
			tf.offset = 0
		}
	}

	for i := 0; i < width; i++ {
		runeIdx := tf.offset + i
		fg := tf.Foreground
		bg := tf.Background
		ch := ' '

		if runeIdx < len(runes) {
			ch = runes[runeIdx]
		} else if !tf.focused && tf.Value == "" && runeIdx < len([]rune(tf.Placeholder)) {
			ch = []rune(tf.Placeholder)[runeIdx]
			fg = tf.PlaceholderForeground
		}

		if runeIdx == tf.cursor && tf.focused {
			fg = tf.CursorForeground
			bg = tf.CursorBackground
		}

		ctx.Draw(i, 0, renderer.Cell{
			Ch:         ch,
			Foreground: fg,
			Background: bg,
		})
	}
}
