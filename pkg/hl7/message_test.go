package hl7

import (
	"fmt"
	"testing"
	"time"

	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7/enums"
)

func TestCreateAck(t *testing.T) {

	messageStr := "MSH|^~\\&|PAS-A|HOS-A|PAS-B|HOS-B|198808181126|SECURITY|ADT^A01^ADT_A01|MSG00001|P|2.8||"
	msg := ParseMessage(messageStr)

	now := time.Now()
	ack := msg.CreateAckMessage()

	hl7Time := now.Format("20060102150405")
	expected := fmt.Sprintf("MSH|^~\\&|PAS-B|HOS-B|PAS-A|HOS-A|%s||ACK|ACK%s|P|2.8\rMSA|AA|MSG00001", hl7Time, hl7Time)
	actual := ack.ToString(StandardDelimters())

	if actual != expected {
		t.Errorf("got %s but wanted %s", actual, expected)
	}
}

func TestCreateNack(t *testing.T) {

	messageStr := "MSH|^~\\&|PAS-A|HOS-A|PAS-B|HOS-B|198808181126|SECURITY|ADT^A01^ADT_A01|MSG00001|P|2.8||"
	msg := ParseMessage(messageStr)

	now := time.Now()
	ack := msg.CreateNackMessage(enums.ApplicationReject)

	hl7Time := now.Format("20060102150405")
	expected := fmt.Sprintf("MSH|^~\\&|PAS-B|HOS-B|PAS-A|HOS-A|%s||ACK|ACK%s|P|2.8\rMSA|AR|MSG00001", hl7Time, hl7Time)
	actual := ack.ToString(StandardDelimters())

	if actual != expected {
		t.Errorf("got %s but wanted %s", actual, expected)
	}
}
