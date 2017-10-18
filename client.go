package network

import (
	"fmt"
	"net"
	"time"
)

type TcpClient interface {
	Say(message string) (string, error)
	Cleanup() error
}

type tcpClient struct {
	conn net.Conn
}

func NewTcpClient(
	address string,
	timeout time.Duration,
) (TcpClient, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &tcpClient{
		conn: conn,
	}, nil
}

func (client tcpClient) Say(
	message string,
) (string, error) {
	return "", nil
}

func (client tcpClient) Cleanup() error {
	return client.conn.Close()
}

// RunTcpClient -
// Starts a client and blocks on stdin forever,
// waiting on bytes and writes them immediately to
// the connection
func RunTcpClient(
	address string,
	timeout time.Duration,
) (err error) {
	client, err := NewTcpClient(address, timeout)
	if err != nil {
		return
	}
	defer func() {
		if err == nil {
			// if we did not set the error, an interrupt or panic happened
			fmt.Println("Interrupt/panic!")
		} else {
			fmt.Println(err.Error())
		}
		fmt.Println("Closing connection")
		err = client.Cleanup()
	}()

	for line, err := NextStdioLine(); err != nil; line, err = NextStdioLine() {
		client.Say(line)
	}
	return
}
