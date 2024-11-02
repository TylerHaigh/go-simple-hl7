package server

import (
	hl7 "github.com/TylerHaigh/go-simple-hl7/pkg/hl7"
)

type Req struct {
	Message hl7.Message
}
