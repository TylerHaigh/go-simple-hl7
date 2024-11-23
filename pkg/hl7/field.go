package hl7

import "strings"

type ComponentString string

//type FieldComponentList []ComponentString

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

func ParseFieldPointer(s string) *Field {
	f := ParseField(s)
	return &f
}

func NewField(c []*Component) Field {
	f := Field{
		Components: c,
	}

	return f
}

func NewFieldPointer(c []*Component) *Field {
	f := NewField(c)
	return &f
}

func FieldFromComponents(componentStrings []ComponentString) *Field {
	components := []*Component{}

	for _, c := range componentStrings {
		component := ParseComponent(string(c))
		components = append(components, &component)
	}

	field := NewField(components)
	return &field
}

func (f *Field) ToString(d Delimeters) string {

	str := ""
	fieldLen := len(f.Components)

	for i, c := range f.Components {
		str += c.ToString(d)

		if i != fieldLen-1 {
			str += string(d.ComponentSeparator)
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
