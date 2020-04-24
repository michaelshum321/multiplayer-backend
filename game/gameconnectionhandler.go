package game

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type GameConnectionHandlerI interface {
	ParseContent([]byte) Command
}

type GameConnectionHandlerS struct {
	world *World
}

type GameConnectionHandler struct {
	GameConnectionHandlerS
}

func (g GameConnectionHandler) HandleConnection(c websocket.Conn) {
	log.Println("Serving", c.RemoteAddr().String())
	defer c.Close()
	for {
		_, message, err := c.ReadMessage()
		//netData, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Println("Error reading from connection", err)
			return
		}
		log.Println("received", string(message), "from", c.RemoteAddr().String())
		g.world.AddCommand(g.ParseContent(message))
		c.WriteMessage(websocket.TextMessage, []byte{1})
		//c.Write([]byte("1"))
	}
}

// {"ModelId": "123", "Dir": 3}
func (g GameConnectionHandler) ParseContent(in []byte) Command {
	log.Println("parsing", string(in))
	var command Command
	err := json.Unmarshal(in, &command)
	if err != nil {
		log.Println("could not unmarshal", err)
	}
	return command
}

func NewGameConnectionHandler(world *World) GameConnectionHandler {
	return GameConnectionHandler{
		GameConnectionHandlerS: GameConnectionHandlerS{
			world: world,
		},
	}
}
