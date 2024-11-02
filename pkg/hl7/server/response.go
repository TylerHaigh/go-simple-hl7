package server

import (
	"net"

	hl7 "github.com/TylerHaigh/go-simple-hl7/pkg/hl7"
)

type Res struct {
	Ack  hl7.Message
	Conn net.Conn
}

func (r *Res) End() {
	ackStringBytes := ([]byte)(r.Ack.ToString(hl7.StandardDelimters()))
	var bytes []byte

	bytes = append(bytes, VT)
	bytes = append(bytes, ackStringBytes...)
	bytes = append(bytes, FS)
	bytes = append(bytes, CR)

	r.Conn.Write(bytes)
}
