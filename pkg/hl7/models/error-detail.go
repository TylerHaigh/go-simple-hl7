package models

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

func NewErrorDetail(
	e error,
	code enums.MessageErrorConditionCodes,
	severity enums.ErrorSeverity,
) *ErrorDetail {

	err := ErrorDetail{
		Location:    nil,
		Code:        enums.ApplicationInternalError,
		Severity:    enums.Error,
		Description: e.Error(),
	}

	return &err
}
