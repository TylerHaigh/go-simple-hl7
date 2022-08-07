package hl7

import "strings"

type Component struct {
	Data []string
}

func NewComponent(s []string) Component {
	c := Component{
		Data: s,
	}
	return c
}

func ParseComponent(s string) Component {
	subcomponents := strings.Split(s, string(StandardDelimters().SubComponentSeparator))
	c := Component{
		Data: subcomponents,
	}
	return c
}

func (c *Component) ToString(d Delimeters) string {

	str := ""
	compLen := len(c.Data)

	for i, sc := range c.Data {
		str += sc
		if i != compLen {
			str += string(d.SubComponentSeparator)
		}
	}

	return str
}

func (c *Component) GetSubComponent(idx uint) string {
	if int(idx) > len(c.Data) {
		return ""
	}

	comp := c.Data[idx-1]
	return comp
}
