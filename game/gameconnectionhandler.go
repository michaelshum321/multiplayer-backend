package game

import (
	"bufio"
	"encoding/json"
	"log"
	"net"
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

func (g GameConnectionHandler) HandleConnection(c net.Conn) {
	log.Println("Serving", c.RemoteAddr().String())
	defer c.Close()
	for {
		netData, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Println("Error reading from connection", err)
			return
		}
		log.Println("received", string(netData), "from", c.RemoteAddr().String())
		g.world.AddCommand(g.ParseContent(netData))
		c.Write([]byte("1"))
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
