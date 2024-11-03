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
	bytes := WrapInEnvelope((r.Ack))
	r.Conn.Write(bytes)
}
