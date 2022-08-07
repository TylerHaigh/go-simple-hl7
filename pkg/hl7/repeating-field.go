package hl7

import "strings"

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

	rpt := RepeatingField{
		Repeats: fields,
	}

	return rpt
}

func NewRepeatingField(fields []*Field) RepeatingField {
	rpt := RepeatingField{
		Repeats: fields,
	}

	return rpt
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

func (r *RepeatingField) GetSubComponent(repeatIndex uint, componentIndex uint, subComponentIndex uint) string {
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
		if i != lenField {
			str += string(d.RepeatSeparator)
		}
	}

	return str

}
