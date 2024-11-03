package messaging

import "github.com/TylerHaigh/go-simple-hl7/pkg/hl7"

const (
	VT = byte(0x0b)
	FS = byte(0x1c)
	CR = byte(0x0d)
)

type Hl7Message string

func WrapInEnvelope(msg hl7.Message) []byte {
	msgString := msg.ToString(hl7.StandardDelimters())
	return WrapStringInEnvelope(Hl7Message(msgString))
}

func WrapStringInEnvelope(msg Hl7Message) []byte {

	msgBytes := []byte(msg)

	var bytes []byte
	bytes = append(bytes, VT)
	bytes = append(bytes, msgBytes...)
	bytes = append(bytes, FS)
	bytes = append(bytes, CR)

	return bytes
}
