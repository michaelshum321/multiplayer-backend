package game

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"multiplayer-backend/game/entity"
	"strconv"
)

type GameConnectionHandlerI interface {
	ParseContent([]byte) Command
}

type GameConnectionHandlerS struct {
	world *World
	// TODO: instead of using IP, use some form of session ID
	// so highjacking is less possible
	connections        map[string]int          // ip -> playerId
	entityToConnection map[int]*websocket.Conn // playerId -> connection
}

type GameConnectionHandler struct {
	GameConnectionHandlerS
}

/**
const (
	YOUR_NEW_PLAYER = "YOUR_NEW_PLAYER" // your new player's new location
	MOVE = "MOVE" // other player's (new) location
)
type Response struct{
	command string
	payload map[string]interface
}
*/
type PlayerIdResponse struct {
	PlayerID int
	X        entity.GridType
	Y        entity.GridType
}

func (g GameConnectionHandler) newPlayer(c *websocket.Conn, connectionId string) {
	startPos := entity.GridType(len(g.connections) * 2)
	newPlayerId := g.world.NewPerson(startPos, startPos)
	response := NewPlayerIdResponse(newPlayerId, startPos, startPos)
	// first msg to new player is their location
	c.WriteJSON(response)
	// send existing connections new player's location
	for otherPlayerId := range g.entityToConnection {
		otherEntity := g.world.objects[strconv.Itoa(otherPlayerId)]
		otherX, otherY := otherEntity.GetPosition()
		c.WriteJSON(NewPlayerIdResponse(otherPlayerId, otherX, otherY))
	}
	g.connections[connectionId] = newPlayerId
	if len(g.connections) == 1 {
		go g.world.StartTime()
	}

	// send other people's location to new players
	g.entityToConnection[newPlayerId] = c
	g.world.broadcastChannel <- *response
}

func (g GameConnectionHandler) HandleConnection(c *websocket.Conn) {
	connectionId := c.RemoteAddr().String()
	log.Println("Serving", connectionId)
	defer c.Close()
	for {
		// TODO: handle game logic for new players elsewhere?
		playerId, ok := g.connections[connectionId]
		if !ok {
			g.newPlayer(c, connectionId)
			continue
		}
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Error reading from connection", err)
			return
		}
		log.Println("received", string(message), "from", c.RemoteAddr().String())
		cmd := g.ParseContent(message)
		cmd.ModelId = strconv.Itoa(playerId)
		g.world.AddCommand(cmd)
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

func NewGameConnectionHandler(world *World) (g GameConnectionHandler) {
	g = GameConnectionHandler{
		GameConnectionHandlerS: GameConnectionHandlerS{
			world:              world,
			connections:        make(map[string]int),
			entityToConnection: make(map[int]*websocket.Conn),
		},
	}
	go g.responsesLoop() // start loop to send responses when events added to channel
	return g
}

func (g GameConnectionHandler) responsesLoop() {
	for {
		select {
		case resp, ok := <-g.world.broadcastChannel:
			log.Print("new response", resp)
			if ok {
				// write to all conn's except self
				for id, conn := range g.entityToConnection {
					if id == resp.PlayerID {
						continue
					}
					log.Println("sending response to id ", id)
					conn.WriteJSON(&resp)
				}

			} else {
				log.Println("responses channel dead")
			}
		}
	}
}

func NewPlayerIdResponse(playerId int, x entity.GridType, y entity.GridType) *PlayerIdResponse {

	return &PlayerIdResponse{
		PlayerID: playerId,
		X:        x,
		Y:        y,
	}
}
