package hl7

import "strings"

type SubComponent string

type Component struct {
	Data []SubComponent
}

func NewComponent(s []SubComponent) Component {
	c := Component{
		Data: s,
	}
	return c
}

func ParseComponent(s string) Component {
	subcomponentStrings := strings.Split(s, string(StandardDelimters().SubComponentSeparator))

	// covert strings to SubComponents
	subcomponents := make([]SubComponent, len(subcomponentStrings))
	for i := range subcomponentStrings {
		subcomponents[i] = SubComponent(subcomponentStrings[i])
	}

	return NewComponent(subcomponents)
}

func (c *Component) ToString(d Delimeters) string {

	str := ""
	compLen := len(c.Data)

	for i, sc := range c.Data {
		str += string(sc)
		if i != compLen {
			str += string(d.SubComponentSeparator)
		}
	}

	return str
}

func (c *Component) GetSubComponent(idx uint) SubComponent {
	if int(idx) > len(c.Data) {
		return ""
	}

	comp := c.Data[idx-1]
	return comp
}
