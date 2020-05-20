package game

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

type GameConnectionHandlerI interface {
	ParseContent([]byte) Command
}

type GameConnectionHandlerS struct {
	world *World
	// TODO: instead of using IP, use some form of session ID
	// so highjacking is less possible
	connections map[string]int
}

type GameConnectionHandler struct {
	GameConnectionHandlerS
}

type PlayerIdResponse struct {
	PlayerID int
}

func (g GameConnectionHandler) HandleConnection(c websocket.Conn) {
	connectionId := c.RemoteAddr().String()
	log.Println("Serving", connectionId)
	defer c.Close()
	for {
		playerId, ok := g.connections[connectionId]
		if !ok {
			log.Println("new cx in gameConnHandler", connectionId)
			newPlayerId := g.world.NewPerson(2, 2)
			playerIdJson, _ := json.Marshal(&PlayerIdResponse{PlayerID: newPlayerId})
			c.WriteMessage(
				websocket.TextMessage,
				playerIdJson,
			)
			g.connections[connectionId] = newPlayerId
			if len(g.connections) == 1 {
				go g.world.StartTime()
			}
			continue
		}
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading from connection", err)
			return
		}
		log.Println("received", string(message), "from", c.RemoteAddr().String())
		cmd := g.ParseContent(message)
		// TODO: remove modelId from front-end response & Command here
		cmd.ModelId = strconv.Itoa(playerId)
		g.world.AddCommand(cmd)
		c.WriteMessage(websocket.TextMessage, []byte(":D"))
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
			world:       world,
			connections: make(map[string]int),
		},
	}
}
