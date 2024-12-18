package hl7

import (
	"strings"
)

type RepeatingFieldList [][]ComponentString

type RepeatingField struct {
	Repeats []*Field
}

func ParseRepeatingField(s string) RepeatingField {
	fieldStrings := strings.Split(s, string(StandardDelimters().RepeatSeparator))
	fields := []*Field{}

	for _, f := range fieldStrings {
		field := ParseField(f)
		fields = append(fields, &field)
	}

	return NewRepeatingField(fields)
}

func ParseRepeatingFieldPointer(s string) *RepeatingField {
	rpt := ParseRepeatingField(s)
	return &rpt
}

func RepeatingFieldFromComponents(fieldArray RepeatingFieldList) *RepeatingField {

	fields := []*Field{}
	for _, components := range fieldArray {
		field := FieldFromComponents(components)
		fields = append(fields, field)
	}

	rpt := NewRepeatingField(fields)
	return &rpt
}

func NewFieldRepeatList(componentStrings []string) RepeatingFieldList {
	components := make([]ComponentString, len(componentStrings))

	for i, s := range componentStrings {
		components[i] = ComponentString(s)
	}

	rpt := RepeatingFieldList{components}
	return rpt
}

func NewRepeatingField(fields []*Field) RepeatingField {
	rpt := RepeatingField{
		Repeats: fields,
	}

	return rpt
}

func NewRepeatingFieldPointer(fields []*Field) *RepeatingField {
	rpt := NewRepeatingField(fields)
	return &rpt
}

func (r *RepeatingField) GetRepeat(repeatIndex uint) *Field {
	if int(repeatIndex) > len(r.Repeats) {
		return nil
	}

	return r.Repeats[repeatIndex-1]
}

func (r *RepeatingField) GetComponent(repeatIndex uint, componentIndex uint) *Component {
	rpt := r.GetRepeat(repeatIndex)
	if rpt == nil {
		return nil
	}

	return rpt.GetComponent(componentIndex)
}

func (r *RepeatingField) GetSubComponent(repeatIndex uint, componentIndex uint, subComponentIndex uint) SubComponent {
	comp := r.GetComponent(repeatIndex, componentIndex)
	if comp != nil {
		return ""
	}

	return comp.GetSubComponent(subComponentIndex)
}

func (r *RepeatingField) GetFieldString(repeatIndex uint) string {
	f := r.GetRepeat(repeatIndex)
	if f == nil {
		return ""
	}

	return f.ToString(StandardDelimters())
}

func (r *RepeatingField) GetComponentString(repeatIndex uint, componentIndex uint) string {
	comp := r.GetComponent(repeatIndex, componentIndex)
	if comp == nil {
		return ""
	}

	return comp.ToString(StandardDelimters())
}

func (r *RepeatingField) ToString(d Delimeters) string {

	str := ""
	lenField := len(r.Repeats)

	for i, f := range r.Repeats {
		str += f.ToString(d)
		if i != lenField-1 {
			str += string(d.RepeatSeparator)
		}
	}

	return str

}
