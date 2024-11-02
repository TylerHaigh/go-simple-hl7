package server

import (
	"bufio"
	"bytes"
	"fmt"
	"net"

	hl7 "github.com/TylerHaigh/go-simple-hl7/pkg/hl7"
)

type SimpleHl7TcpServer struct {
	listener      net.Listener
	handlers      []HandlerFunction
	errorHandlers []ErrorFunction
}

type ConnectionDetails struct {
	Host string
	Port string
}

func (s *SimpleHl7TcpServer) Start(conn ConnectionDetails) error {
	address := fmt.Sprintf("%s:%s", conn.Host, conn.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	s.listener = listener

	return nil
}

func (s *SimpleHl7TcpServer) Close() {
	s.listener.Close()
}

func (s *SimpleHl7TcpServer) AcceptConnection() {
	c, err := s.listener.Accept()
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	fmt.Println("Client connected.")

	fmt.Println("Client " + c.RemoteAddr().String() + " connected.")

	go s.handleConnection(c)
}

func (s *SimpleHl7TcpServer) handleConnection(conn net.Conn) {

	defer conn.Close()
	messageBuffer := bytes.NewBuffer([]byte{})

	buffer, err := bufio.NewReader(conn).ReadBytes(FS)

	if err != nil {
		fmt.Println("Client left.")
		conn.Close()
		return
	}

	messageBuffer.Write(buffer)

	vt := make([]byte, 1)
	messageBuffer.Read(vt)
	messageBuffer.Truncate(messageBuffer.Len() - 1)
	messageStr, _ := messageBuffer.ReadString(FS)

	message := hl7.ParseMessage(messageStr)
	req := Req{Message: message}
	res := Res{Ack: message.CreateAckMessage(), Conn: conn}

	ctx := DefaultCtx{
		Req:               &req,
		Res:               &res,
		handlerIndex:      0,
		handlers:          s.handlers,
		errorHandlerIndex: 0,
		errorHandlers:     s.errorHandlers,
	}

	s.handleRequest(&ctx)
}

func (s *SimpleHl7TcpServer) Use(fn HandlerFunction) {
	s.handlers = append(s.handlers, fn)
}

func (s *SimpleHl7TcpServer) HandleError(fn ErrorFunction) {
	s.errorHandlers = append(s.errorHandlers, fn)
}

func (s *SimpleHl7TcpServer) handleRequest(ctx *DefaultCtx) error {

	ctx.handlerIndex = 0
	ctx.errorHandlerIndex = 0
	stackSize := len(ctx.handlers)

	if stackSize > 0 {
		// Get the first Handler
		h := ctx.handlers[ctx.handlerIndex]

		// Execute the Handler
		err := h(ctx)

		if err != nil {
			ctx.Error = err
		}

	}

	return ctx.Error

}
