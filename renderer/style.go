package renderer

type Style struct {
	Bold      bool
	Italic    bool
	Underline bool
}

func (s Style) Sequence() string {
	seq := "\x1b[0m" // reset sempre
	if s.Bold {
		seq += "\x1b[1m"
	}
	if s.Italic {
		seq += "\x1b[3m"
	}
	if s.Underline {
		seq += "\x1b[4m"
	}
	return seq
}
