package server

import (
	"net"

	hl7 "github.com/TylerHaigh/go-simple-hl7/pkg/hl7"
	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7/messaging"
)

type Res struct {
	Ack  hl7.Message
	Conn net.Conn
}

func (r *Res) End() {
	bytes := messaging.WrapInEnvelope((r.Ack))
	r.Conn.Write(bytes)
}
