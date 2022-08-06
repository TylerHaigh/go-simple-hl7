package hl7

type Delimeters struct {
	MessageSeparator      rune // \n
	SegmentSeparator      rune // \r
	FieldSeparator        rune // |
	RepeatSeparator       rune // ~
	ComponentSeparator    rune // ^
	SubComponentSeparator rune // &
}

func StandardDelimters() Delimeters {
	d := Delimeters{
		MessageSeparator:      '\n',
		SegmentSeparator:      '\r',
		FieldSeparator:        '|',
		RepeatSeparator:       '~',
		ComponentSeparator:    '^',
		SubComponentSeparator: '&',
	}

	return d
}
