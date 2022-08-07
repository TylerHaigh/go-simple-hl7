package hl7

import "strings"

type Field struct {
	Components []*Component
}

func ParseField(s string) Field {

	componentStrings := strings.Split(s, string(StandardDelimters().ComponentSeparator))
	components := []*Component{}

	for _, c := range componentStrings {
		comp := ParseComponent(c)
		components = append(components, &comp)
	}

	f := Field{
		Components: components,
	}

	return f
}

func NewField(c []*Component) Field {
	f := Field{
		Components: c,
	}

	return f
}

func (f *Field) ToString(d Delimeters) string {

	str := ""
	fieldLen := len(f.Components)

	for i, c := range f.Components {
		str += c.ToString(d)

		if i != fieldLen {
			str += string(d.FieldSeparator)
		}
	}

	return str

}

func (f *Field) GetComponentString(componentIndex uint) string {
	comp := f.GetComponent(componentIndex)
	if comp != nil {
		return comp.ToString(StandardDelimters())
	}

	return ""
}

func (f *Field) GetComponent(componentIndex uint) *Component {
	if int(componentIndex) > len(f.Components) {
		return nil
	}

	return f.Components[componentIndex-1]
}

func (f *Field) GetSubComponent(componentIndex uint, subComponentIndex uint) SubComponent {
	comp := f.GetComponent(componentIndex)
	if comp == nil {
		return ""
	}

	return comp.GetSubComponent(subComponentIndex)

}
