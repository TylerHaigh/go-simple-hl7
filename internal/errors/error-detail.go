package errors

import "github.com/TylerHaigh/go-simple-hl7/pkg/hl7/enums"

type ErrorLocation struct {
	Segment         string
	SegmentSequence int
	Field           int
	Repetition      int
	Component       int
	SubComponent    int
}

type ErrorDetail struct {
	Location    *ErrorLocation
	Code        enums.MessageErrorConditionCodes
	Severity    enums.ErrorSeverity
	Description string
}
