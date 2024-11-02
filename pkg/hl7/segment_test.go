package hl7

import (
	"testing"
)

func TestSegmentToString_MSH(t *testing.T) {
	originalString := "MSH|^~\\&|ADT1|GOOD HEALTH HOSPITAL"
	seg := ParseSegment(originalString)
	str := seg.ToString(StandardDelimters())
	if str != originalString {
		t.Errorf("got %s but wanted %s", str, originalString)
	}

}

func TestSegmentToString_PID(t *testing.T) {
	originalString := "PID|1||PATID1234||EVERYMAN^ADAM^|"
	seg := ParseSegment(originalString)
	str := seg.ToString(StandardDelimters())
	if str != originalString {
		t.Errorf("got %s but wanted %s", str, originalString)
	}

}
