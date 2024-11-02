package hl7

import (
	"fmt"
	"strings"
	"time"

	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7/enums"
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

func ParseMessagePointer(s string) *Message {
	m := ParseMessage(s)
	return &m
}

func NewMessage(segments []*Segment) Message {
	msh := Message{
		Segments: segments,
	}

	return msh
}

func NewMessagePointer(segments []*Segment) *Message {
	m := NewMessage(segments)
	return &m
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
	lenSegments := len(m.Segments)

	for i, s := range m.Segments {
		str += s.ToString(d)
		if i != lenSegments-1 {
			str += string(d.SegmentSeparator)
		}
	}

	return str
}

func (m *Message) CreateAckMessage() Message {

	t := time.Now()

	// https://www.geeksforgeeks.org/time-formatting-in-golang/
	// https://www.practical-go-lessons.com/post/how-to-format-time-with-golang-ccc5ja83ibmc70m98260

	hl7Time := t.Format("20060102150405")

	msh := NewSegment("MSH", []*RepeatingField{
		ParseRepeatingFieldPointer("|"),     // MSH-1
		ParseRepeatingFieldPointer("^~\\&"), // MSH-2
		m.GetField("MSH", 5),                // MSH-3: Sending application (obtain from current receiving application MSH-5)
		m.GetField("MSH", 6),                // MSH-4: Sending facility (obtain from current receiving facility MSH-6)
		m.GetField("MSH", 3),                // MSH-5: Receiving application (obtain current from sending application MSH-3)
		m.GetField("MSH", 4),                // MSH-6: Receiving facility (obtain from current sending facility MSH-4)
		ParseRepeatingFieldPointer(hl7Time), // MSH-7: Date/Time
		ParseRepeatingFieldPointer(""),      // MSH-8: Security (leave blank)
		ParseRepeatingFieldPointer("ACK"),   // MSH-9: MessageType (ACK)
		ParseRepeatingFieldPointer(
			fmt.Sprintf("ACK%s", hl7Time), // MSH-10: Message Control ID (generated)
		),
		m.GetField("MSH", 11), // MSH-11: Processing ID
		m.GetField("MSH", 12), // MSH-12: Version ID
	})

	msa := NewSegment("MSA", []*RepeatingField{
		ParseRepeatingFieldPointer("AA"),
		m.GetField("MSH", 10),
	})

	segments := []*Segment{&msh, &msa}
	ack := NewMessage(segments)

	return ack
}

func (m *Message) CreateNackMessage(acknowledgementCode enums.AcknowledgementCode) Message {
	t := time.Now()

	// https://www.geeksforgeeks.org/time-formatting-in-golang/
	// https://www.practical-go-lessons.com/post/how-to-format-time-with-golang-ccc5ja83ibmc70m98260
	//
	// Whereas other languages use a format like YYYY-MM-DD to format dates like: 2022-10-21,
	// Go uses a reference time.
	// This reference time is a point in time that the language will use to parse your
	// layout : - 2 January 2006 03:04:05 PM in the time zone UTC -7
	//
	// You might ask why this particular date. That’s because when you read it like that :
	// 01/02 03:04:05PM '06 -0700
	//
	// You can note that numbers follow each other: 1 (January), 2 (day), 3 (hour), 4(minutes)…

	hl7Time := t.Format("20060102150405")

	msh := NewSegment("MSH", []*RepeatingField{
		ParseRepeatingFieldPointer("|"),     // MSH-1
		ParseRepeatingFieldPointer("^~\\&"), // MSH-2
		m.GetField("MSH", 5),                // MSH-3: Sending application (obtain from current receiving application MSH-5)
		m.GetField("MSH", 6),                // MSH-4: Sending facility (obtain from current receiving facility MSH-6)
		m.GetField("MSH", 3),                // MSH-5: Receiving application (obtain current from sending application MSH-3)
		m.GetField("MSH", 4),                // MSH-6: Receiving facility (obtain from current sending facility MSH-4)
		ParseRepeatingFieldPointer(hl7Time), // MSH-7: Date/Time
		ParseRepeatingFieldPointer(""),      // MSH-8: Security (leave blank)
		ParseRepeatingFieldPointer("ACK"),   // MSH-9: MessageType (ACK)
		ParseRepeatingFieldPointer(
			fmt.Sprintf("ACK%s", hl7Time), // MSH-10: Message Control ID (generated)
		),
		m.GetField("MSH", 11), // MSH-11: Processing ID
		m.GetField("MSH", 12), // MSH-12: Version ID
	})

	msa := NewSegment("MSA", []*RepeatingField{
		ParseRepeatingFieldPointer(acknowledgementCode.String()),
		m.GetField("MSH", 10),
	})

	// TODO: Create ERR segments from an error listing

	//     const version = nack.getHeader().versionId.toString();
	//     for (const e of errors) {
	//       const err = createErrorSegment(version, e);
	//       nack.pushSegment(err);
	//     }

	segments := []*Segment{&msh, &msa}
	nack := NewMessage(segments)

	return nack
}
