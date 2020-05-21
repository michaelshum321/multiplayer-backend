package game

import "log"

type Actions struct {
	queue chan Command
}

type Direction uint8

const (
	Down  Direction = 0
	Up    Direction = 1
	Left  Direction = 2
	Right Direction = 3
)

// for now, just do 'Move'
type Command struct {
	ModelId string
	Direction     Direction
}

func (actions *Actions) addCommand(command Command) {
	actions.queue <- command
	log.Println("added command ", command)
}
