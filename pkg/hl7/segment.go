package hl7

import (
	"fmt"
	"strings"
)

type Segment struct {
	Name   string
	Fields []*RepeatingField
}

func ParseSegment(s string) Segment {
	fieldStrings := strings.Split(s, string(StandardDelimters().FieldSeparator))
	fields := []*RepeatingField{}

	name := fieldStrings[0]

	// Handle case where MSH-1 is the field Separator delimiter
	if name == "MSH" {
		msh1 := ParseRepeatingField("|")
		fields = append(fields, &msh1)
	}

	for _, f := range fieldStrings[1:] {
		field := ParseRepeatingField(f)
		fields = append(fields, &field)
	}

	return NewSegment(name, fields)
}

func ParseSegmentPointer(s string) *Segment {
	seg := ParseSegment(s)
	return &seg
}

func NewSegment(name string, fields []*RepeatingField) Segment {
	segment := Segment{
		Name:   name,
		Fields: fields,
	}

	return segment
}

func NewSegmentPointer(name string, fields []*RepeatingField) *Segment {
	seg := NewSegment(name, fields)
	return &seg
}

func SegmentFromComponentString(name string, fieldArray []RepeatingFieldList) Segment {

	fields := []*RepeatingField{}
	for _, fa := range fieldArray {
		field := RepeatingFieldFromComponents(fa)
		fields = append(fields, &field)
	}

	return NewSegment(name, fields)
}

func (s *Segment) GetField(fieldIndex uint) *RepeatingField {
	if int(fieldIndex) > len(s.Fields) {
		return nil
	}

	return s.Fields[fieldIndex-1]
}

func (s *Segment) GetFieldRepeat(fieldIndex uint, repeatIndex uint) *Field {
	f := s.GetField(fieldIndex)
	if f == nil {
		return nil
	}

	return f.GetRepeat(repeatIndex)
}

func (s *Segment) GetComponent(fieldIndex uint, repeatIndex uint, componentIndex uint) *Component {
	rpt := s.GetFieldRepeat(fieldIndex, repeatIndex)
	if rpt == nil {
		return nil
	}

	return rpt.GetComponent(componentIndex)
}

func (s *Segment) GetSubComponent(fieldIndex uint, repeatIndex uint, componentIndex uint, subComponentIndex uint) SubComponent {
	c := s.GetComponent(fieldIndex, repeatIndex, componentIndex)
	if c == nil {
		return ""
	}

	return c.GetSubComponent(subComponentIndex)
}

func (s *Segment) GetFieldString(fieldIndex uint) string {
	f := s.GetField(fieldIndex)
	if f == nil {
		return ""
	}

	return f.ToString(StandardDelimters())
}

func (s *Segment) GetFieldRepeatString(fieldIndex uint, repeatIndex uint) string {
	rpt := s.GetFieldRepeat(fieldIndex, repeatIndex)
	if rpt == nil {
		return ""
	}

	return rpt.ToString(StandardDelimters())
}

func (s *Segment) GetComponentString(fieldIndex uint, repeatIndex uint, componentIndex uint) string {
	c := s.GetComponent(fieldIndex, repeatIndex, componentIndex)
	if c == nil {
		return ""
	}

	return c.ToString(StandardDelimters())
}

func (s *Segment) ToString(d Delimeters) string {

	str := fmt.Sprintf("%s%s", s.Name, string(d.FieldSeparator))
	lenFields := len(s.Fields)

	for i, f := range s.Fields {

		// Skip over MSH-1 as this is covered above
		if s.Name == "MSH" && i == 0 {
			continue
		}

		str += f.ToString(d)
		if i != lenFields-1 {
			str += string(d.FieldSeparator)
		}
	}

	return str
}
