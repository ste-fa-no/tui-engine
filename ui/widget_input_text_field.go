package ui

import (
	"tui-engine/events"
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

func (tf *TextField) HandleEvent(e events.Event) bool {
	switch ev := e.(type) {
	case events.KeyEvent:
		switch ev.Code {
		case events.KeyRune:
			runes := []rune(tf.Value)
			runes = append(runes[:tf.cursor], append([]rune{ev.Rune}, runes[tf.cursor:]...)...)
			tf.Value = string(runes)
			tf.cursor++
		case events.KeyBackspace:
			if tf.cursor > 0 {
				runes := []rune(tf.Value)
				runes = append(runes[:tf.cursor-1], runes[tf.cursor:]...)
				tf.Value = string(runes)
				tf.cursor--
			}
		case events.KeyLeft:
			if tf.cursor > 0 {
				tf.cursor--
			}
		case events.KeyRight:
			if tf.cursor < len([]rune(tf.Value)) {
				tf.cursor++
			}
		}

		return true
	}

	return false
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
