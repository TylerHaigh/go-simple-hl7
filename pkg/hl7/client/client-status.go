package client

type TcpClientState int

const (
	Stopped = iota

	/** Socket is connecting, waiting for 'connect' event. */
	Connecting

	/** Socket connection lost, waiting for RECONNECT_INTERVAL until attempting another connect. */
	ReconnectWait

	/** Socket is connected, ready to start sending messages. */
	Ready

	/** Message sent, waiting for acknowledgement. */
	AcknowledgementWait

	/** Socket is closing, waiting for 'close' event. */
	Closing
)
