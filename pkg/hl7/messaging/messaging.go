package messaging

import (
	"bufio"
	"bytes"

	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7"
)

// ReadHl7Message reads a bufio.Reader stream to read the entire HL7 message.
// A standard HL7 message is composed using an envelope including:
// VT + {the message} + FS + CR.
// This function will read the input stream and return the parsed HL7
// Message structure.
func ReadHl7Message(reader *bufio.Reader) (*hl7.Message, error) {
	messageStr, err := ReadHl7MessageString(reader)

	if err != nil {
		return nil, err
	}

	message := hl7.ParseMessage(messageStr)
	return &message, nil
}

func ReadHl7MessageString(reader *bufio.Reader) (string, error) {
	messageBuffer := bytes.NewBuffer([]byte{})

	// Read up until the HL7 Envelope FS character
	buffer, err := reader.ReadBytes(FS)
	if err != nil {
		return "", err
	}

	// Transfer to a separate buffer so we can read from it
	// instead of the source buffer
	messageBuffer.Write(buffer)

	// Read and write the final carriage return CR of the envelope
	// from the input stream
	cr, _ := reader.ReadByte()
	messageBuffer.Write([]byte{cr})

	// Read the VT character from the message buffer
	vt := make([]byte, 1)
	messageBuffer.Read(vt)

	// Read the string data from the buffer up until the FS character
	messageStr, _ := messageBuffer.ReadString(FS)

	// Trim FS character
	messageStr = messageStr[:len(messageStr)-1]
	return messageStr, nil
}
