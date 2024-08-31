package hl7

import (
	"strings"

	"golang.org/x/exp/slices"
)

type Message struct {
	Segments []*Segment
}

func ParseMessage(s string) Message {
	segmentStrings := strings.Split(s, string(StandardDelimters().SegmentSeparator))
	segments := []*Segment{}

	for _, seg := range segmentStrings {
		segment := ParseSegment(seg)
		segments = append(segments, &segment)
	}

	return NewMessage(segments)
}

func NewMessage(segments []*Segment) Message {
	msh := Message{
		Segments: segments,
	}

	return msh
}

func (m *Message) GetSegment(segmentName string) *Segment {
	idx := slices.IndexFunc(m.Segments, func(s *Segment) bool { return s.Name == segmentName })

	if idx == -1 {
		return nil
	}

	return m.Segments[idx]
}

func (m *Message) GetField(segmentName string, fieldIndex uint) *RepeatingField {
	s := m.GetSegment(segmentName)
	if s == nil {
		return nil
	}

	return s.GetField(fieldIndex)
}

func (m *Message) GetFieldRepeat(segmentName string, fieldIndex uint, repeatIndex uint) *Field {
	f := m.GetField(segmentName, fieldIndex)
	if f == nil {
		return nil
	}

	return f.GetRepeat(repeatIndex)
}

func (m *Message) GetComponent(segmentName string, fieldIndex uint, repeatIndex uint, componentIndex uint) *Component {
	f := m.GetFieldRepeat(segmentName, fieldIndex, repeatIndex)
	if f == nil {
		return nil
	}

	return f.GetComponent(componentIndex)
}

func (m *Message) GetSubComponent(
	segmentName string, fieldIndex uint, repeatIndex uint,
	componentIndex uint, subComponentIndex uint,
) SubComponent {

	c := m.GetComponent(segmentName, fieldIndex, repeatIndex, componentIndex)
	if c == nil {
		return ""
	}

	return c.GetSubComponent(subComponentIndex)
}

func (m *Message) GetSegmentString(segmentName string) string {
	s := m.GetSegment(segmentName)
	if s == nil {
		return ""
	}

	return s.ToString(StandardDelimters())
}

func (m *Message) GetFieldString(segmentName string, fieldIndex uint) string {
	f := m.GetField(segmentName, fieldIndex)
	if f == nil {
		return ""
	}

	return f.ToString(StandardDelimters())
}

func (m *Message) GetFieldRepeatString(segmentName string, fieldIndex uint, repeatIndex uint) string {
	f := m.GetFieldRepeat(segmentName, fieldIndex, repeatIndex)
	if f == nil {
		return ""
	}

	return f.ToString(StandardDelimters())
}

func (m *Message) GetComponentString(
	segmentName string, fieldIndex uint, repeatIndex uint,
	componentString uint,
) string {

	c := m.GetComponent(segmentName, fieldIndex, repeatIndex, componentString)
	if c == nil {
		return ""
	}

	return c.ToString(StandardDelimters())
}


func (m *Message) ToString(d Delimeters) string {

	str := ""
	lenSegments := len(s.Segments)

	for i, s := range s.Segments {
		str += s.ToString(d)
		if i != lenSegments-1 {
			str += string(d.SegmentSeparator)
		}
	}

	return str
}
