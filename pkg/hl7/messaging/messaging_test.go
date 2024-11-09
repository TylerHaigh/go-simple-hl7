package messaging

import (
	"bufio"
	"strings"
	"testing"
)

func TestReadHl7MessageString(t *testing.T) {

	str := "Hello"
	env := WrapStringInEnvelope(Hl7Message(str))
	reader := strings.NewReader(string(env))
	bufReader := *bufio.NewReader(reader)

	message, err := ReadHl7MessageString(&bufReader)
	if err != nil {
		t.Errorf("Error reading HL7 message. Error: %v", err)
	}

	if message != str {
		t.Errorf("got %s but wanted %s", str, message)
	}
}
