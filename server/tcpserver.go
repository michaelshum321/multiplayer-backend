package main

import (
	"log"
	"net"
)

type ServerI interface {
	Start()
}

type ConnectionHandlerI interface{
	HandleConnection(conn net.Conn)
}

type ServerS struct {
	connectionHandler ConnectionHandlerI
}

type TcpServer struct {
	ServerS
}

func (t TcpServer) Start() {
	log.Println("starting server")
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println("error accepting",err)
			return
		}
		go t.ServerS.connectionHandler.HandleConnection(c)
	}
}

func NewTcpServer(connHandler ConnectionHandlerI) TcpServer {
	return TcpServer{
		ServerS: ServerS{
			connectionHandler: connHandler,
		},
	}
}