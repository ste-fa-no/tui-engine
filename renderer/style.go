package renderer

type Style struct {
	Bold      bool
	Italic    bool
	Underline bool
}

type Styles int

const (
	Reset Styles = iota
	Bold
	Italic
	Underline
)

var styleSequenceMap = map[Styles]string{
	Reset:     "\x1b[0m",
	Bold:      "\x1b[1m",
	Italic:    "\x1b[3m",
	Underline: "\x1b[4m",
}

func (s Style) Sequence() string {
	seq := styleSequenceMap[Reset]

	if s.Bold {
		seq += styleSequenceMap[Bold]
	}
	if s.Italic {
		seq += styleSequenceMap[Italic]
	}
	if s.Underline {
		seq += styleSequenceMap[Underline]
	}

	return seq
}
