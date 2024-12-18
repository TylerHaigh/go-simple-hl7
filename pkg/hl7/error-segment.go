package hl7

import (
	"fmt"
	"strconv"

	"github.com/TylerHaigh/go-simple-hl7/internal/errors"
	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7/models"
)

const minimumErrorSegmentVersion = 2.5

func valueOrNull(val int) string {
	if val == 0 {
		return ""
	}

	return string(val)
}

func CreateErrorSegment(hl7Version string, errorDetails models.ErrorDetail) *Segment {

	isVersion251 := hl7Version == "2.5.1"
	v, _ := strconv.ParseFloat(hl7Version, 32)
	if isVersion251 || v > minimumErrorSegmentVersion {
		return enhancedErrorSegment(errorDetails)
	} else {
		return legacyErrorSegment(errorDetails)
	}

}

func enhancedErrorSegment(errorDetails models.ErrorDetail) *Segment {
	desc := errors.MapErrorCodeDescription(errorDetails.Code)

	errorLocation := []string{}
	if errorDetails.Location != nil {
		errorLocation = []string{
			errorDetails.Location.Segment,
			string(errorDetails.Location.SegmentSequence),
			valueOrNull(errorDetails.Location.Field),
			valueOrNull(errorDetails.Location.Repetition),
			valueOrNull(errorDetails.Location.Component),
			valueOrNull(errorDetails.Location.SubComponent),
		}
	}

	field := []RepeatingFieldList{
		NewFieldRepeatList([]string{""}),  // ERR-1
		NewFieldRepeatList(errorLocation), // ERR-2
		NewFieldRepeatList([]string{
			errorDetails.Code.String(), desc, "HL70357", // ERR-3
		}), // ERR-3
		NewFieldRepeatList([]string{errorDetails.Severity.String()}), // ERR-4
		NewFieldRepeatList([]string{""}),                             // ERR-5
		NewFieldRepeatList([]string{""}),                             // ERR-6
		NewFieldRepeatList([]string{""}),                             // ERR-7
		NewFieldRepeatList([]string{errorDetails.Description}),       // ERR-8

	}

	err := SegmentFromComponentString("ERR", field)
	return err
}

func legacyErrorSegment(errorDetails models.ErrorDetail) *Segment {

	desc := errors.MapErrorCodeDescription(errorDetails.Code)

	errorLocation := []string{}
	if errorDetails.Location != nil {
		errorLocation = []string{
			errorDetails.Location.Segment,
			string(errorDetails.Location.SegmentSequence),
			valueOrNull(errorDetails.Location.Field),
			fmt.Sprintf(
				"%s&%s&%s&%s",
				errorDetails.Code.String(),
				desc,
				"HL70357",
				errorDetails.Description,
			),
		}
	}

	field := []RepeatingFieldList{
		NewFieldRepeatList(errorLocation), // ERR-1
	}

	err := SegmentFromComponentString("ERR", field)
	return err
}
