package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

//type ServerI interface {
//	Start()
//}

type ServerI interface {
	Start(http.ResponseWriter, *http.Request)
}

type ConnectionHandlerI interface {
	HandleConnection(conn *websocket.Conn)
}

type ServerS struct {
	connectionHandler ConnectionHandlerI
}
type WebsocketServer struct {
	ServerS
}

var upgrader = websocket.Upgrader{} // use default options

func (ws WebsocketServer) Home(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		ws.ServerS.connectionHandler.HandleConnection(c)
	}
}
func (ws WebsocketServer) Start() {
	log.Println("Starting ws server")
	http.HandleFunc("/", ws.Home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func NewWSServer(connHandler ConnectionHandlerI) WebsocketServer {
	return WebsocketServer{
		ServerS: ServerS{
			connectionHandler: connHandler,
		},
	}
}

//type TcpServer struct {
//	ServerS
//}
//
//func (t TcpServer) Start() {
//	log.Println("starting server")
//	l, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer l.Close()
//	for {
//		c, err := l.Accept()
//		if err != nil {
//			log.Println("error accepting",err)
//			return
//		}
//		go t.ServerS.connectionHandler.HandleConnection(c)
//	}
//}

//func NewTcpServer(connHandler ConnectionHandlerI) TcpServer {
//	return TcpServer{
//		ServerS: ServerS{
//			connectionHandler: connHandler,
//		},
//	}
//}
