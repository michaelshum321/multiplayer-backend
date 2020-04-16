package game

type Actions struct {
	queue chan Command
}

type Direction uint8

const (
	Down	Direction = 0
	Up		Direction = 1
	Left	Direction = 2
	Right	Direction = 3
)

// for now, just do 'Move'
type Command struct {
	modelId string
	direction Direction
}

func (actions *Actions) AddCommand(command Command) {
	actions.queue <- command
}

func (actions *Actions) GetQueue() chan Command {
	return actions.queue
}