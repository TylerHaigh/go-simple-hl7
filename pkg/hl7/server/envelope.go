package server

import hl7 "github.com/TylerHaigh/go-simple-hl7/pkg/hl7"

const (
	VT = byte(0x0b)
	FS = byte(0x1c)
	CR = byte(0x0d)
)

func WrapInEnvelope(msg hl7.Message) []byte {

	ackStringBytes := ([]byte)(msg.ToString(hl7.StandardDelimters()))

	var bytes []byte
	bytes = append(bytes, VT)
	bytes = append(bytes, ackStringBytes...)
	bytes = append(bytes, FS)
	bytes = append(bytes, CR)

	return bytes
}
