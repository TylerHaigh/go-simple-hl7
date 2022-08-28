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

func extractSubcomponentStrings(s string) []SubComponent {
	subcomponentStrings := strings.Split(s, string(StandardDelimters().SubComponentSeparator))

	// covert strings to SubComponents
	subcomponents := make([]SubComponent, len(subcomponentStrings))
	for i := range subcomponentStrings {
		subcomponents[i] = SubComponent(subcomponentStrings[i])
	}

	return subcomponents
}

func ParseComponent(s string) Component {
	subcomponents := extractSubcomponentStrings(s)
	return NewComponent(subcomponents)
}

func (c *Component) ToString(d Delimeters) string {

	str := ""
	compLen := len(c.Data)

	for i, sc := range c.Data {
		str += string(sc)
		if i != compLen-1 {
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

func (c *Component) SetFromString(s string) {
	subcomponents := extractSubcomponentStrings(s)
	c.Set(subcomponents)
}

func (c *Component) Set(s []SubComponent) {
	c.Data = s
}
