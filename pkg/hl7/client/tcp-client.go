package client

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/TylerHaigh/go-simple-hl7/pkg/hl7/messaging"
)

type ISimpleHl7Client interface {
	Start(conn ConnectionDetails) error
	StartTLS(conn TLSConnectionDetails, config *tls.Config) error
	Stop() error
	Send(message messaging.Hl7Message) (messaging.Hl7Message, error)
}

type QueuedMessage struct {
	message      messaging.Hl7Message
	responseChan chan AckResponse
}

type AckResponse struct {
	ack messaging.Hl7Message
	err error
}

type SimpleHl7TcpClient struct {
	conn         net.Conn
	state        TcpClientState
	messageQueue chan QueuedMessage
}

type ConnectionDetails struct {
	Host string
	Port string
}

type TLSConnectionDetails struct {
	Host              string
	Port              string
	ServerCertificate string
	ServerKey         string
}

func NewSimpleHl7TcpClient() SimpleHl7TcpClient {
	c := SimpleHl7TcpClient{
		state:        Stopped,
		messageQueue: make(chan QueuedMessage),
	}

	return c
}

func enableKeepAlive(netConn net.Conn) error {
	// Type assert to *net.TCPConn to set TCP-specific options
	tcpConn, ok := netConn.(*net.TCPConn)
	if !ok {
		return errors.New("not a TCP connection")
	}

	// Enable keepalive with default settings
	if err := tcpConn.SetKeepAlive(true); err != nil {
		return err
	}

	// Set keepalive settings period
	if err := tcpConn.SetKeepAlivePeriod(60 * time.Second); err != nil {
		return err
	}

	return nil
}

type DialFn func() (net.Conn, error)

func (c *SimpleHl7TcpClient) start(dialFn DialFn) error {
	if c.state != Stopped {
		return errors.New("HL7 TCP Client is already started")
	}

	c.state = Connecting

	netConn, err := dialFn()
	if err != nil {
		return err
	}

	enableKeepAlive(netConn)
	c.conn = netConn
	c.state = Ready

	go c.loop()

	return nil
}

func (c *SimpleHl7TcpClient) loop() {

	defer c.conn.Close()
	for c.state != Stopped {

		if c.state == Ready {
			messageToSend := <-c.messageQueue
			ack, err := c.writeMessage(messageToSend.message)
			ackResponse := AckResponse{
				ack: ack,
				err: err,
			}
			messageToSend.responseChan <- ackResponse
		}

	}
}

func (c *SimpleHl7TcpClient) writeMessage(message messaging.Hl7Message) (messaging.Hl7Message, error) {
	messageBytes := messaging.WrapStringInEnvelope(messaging.Hl7Message(message))
	_, err := c.conn.Write(messageBytes)

	if err != nil {
		return "", err
	}

	c.state = AcknowledgementWait

	messageBuffer := bytes.NewBuffer([]byte{})
	reader := bufio.NewReader(c.conn)
	ackBytes, err := reader.ReadBytes(messaging.FS)

	if err != nil {
		return "", err
	}

	messageBuffer.Write(ackBytes)
	cr, _ := reader.ReadByte()
	messageBuffer.Write([]byte{cr})

	vt := make([]byte, 1)
	messageBuffer.Read(vt)
	messageBuffer.Truncate(messageBuffer.Len() - 1)
	messageStr, err := messageBuffer.ReadString(messaging.FS)

	// Trim FS character
	messageStr = messageStr[:len(messageStr)-1]

	c.state = Ready

	if err != nil {
		return "", err
	}

	return messaging.Hl7Message(messageStr), nil
}

func (c *SimpleHl7TcpClient) Start(conn TLSConnectionDetails) error {
	address := fmt.Sprintf("%s:%s", conn.Host, conn.Port)

	dialFn := func() (net.Conn, error) {
		netConn, err := net.Dial("tcp", address)
		return netConn, err
	}

	if err := c.start(dialFn); err != nil {
		return err
	}

	return nil
}

func (c *SimpleHl7TcpClient) StartTLS(conn TLSConnectionDetails, config *tls.Config) error {

	address := fmt.Sprintf("%s:%s", conn.Host, conn.Port)

	dialFn := func() (net.Conn, error) {
		netConn, err := tls.Dial("tcp", address, config)
		return netConn, err
	}

	if err := c.start(dialFn); err != nil {
		return err
	}

	return nil
}

func (c *SimpleHl7TcpClient) Send(message messaging.Hl7Message) (messaging.Hl7Message, error) {
	if c.state == Stopped {
		return "", errors.New("cannot send message while mllp connection is stopped")
	}

	responseChan := make(chan AckResponse, 1)

	msg := QueuedMessage{
		message:      message,
		responseChan: responseChan,
	}

	c.messageQueue <- msg

	// Wait for ack to come back
	response := <-responseChan

	return response.ack, response.err

}
