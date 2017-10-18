package network

import (
	"bufio"
	"net"
	"time"
)

type HandlerFn func(req string) (string, error)

type TcpServer interface {
	AddHandler(fn HandlerFn)
	HandleConnection(conn net.Conn)
}

type tcpServer struct {
	Listener net.Listener
	Handlers map[string]HandlerFn
}

func NewTcpServer(address string) (TcpServer, error) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	return &TcpServer{
		Listener: listener,
	}, nil
}

func RunTcpServer(address string) error {
	server, err := NewTcpServer(address)
	if err != nil {
		return err
	}
	server.AddHandler(Poopoo(Rick))
	server.AddHandler(Poopoo(Morty))
	for {
		conn, err := server.Listener.Accept()
		if err != nil {
			fmt.Println(err)
		} else {
			go server.HandleConnection(conn)
		}
	}
}

func HandleConnection(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()

	for {
		select {
		case <-time.After(20 * time.Second):
			fmt.Printf("Connection stale. Closing conn %v\n", conn)
		default:
			request, err := rw.ReadString("\n")
			if err != nil && err != io.EOF {
				fmt.Println(err)
				continue
			}

		}
	}
}

// Handlers
func Rick(req string) (string, error) {
	return fmt.Sprintf("'%s'? I can't even honestly say nice try. Now where the hell is MY Morty, Morty?", req)
}

func Morty(conn net.Conn) (string, error) {
	return "Uhhh oooh I don't think that's a good idea, Rick.."
}
