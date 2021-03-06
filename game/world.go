package game

import (
	"log"
	"multiplayer-backend/game/entity"
	"strconv"
)

type World struct {
	grid             *Grid
	moment           *Timer
	actions          Actions
	objects          map[string]entity.ModelI
	broadcastChannel chan PlayerIdResponse
}

//func (world *World) doSomething() {
//	i := GridType(rand.Uint32()) % world.grid.size
//	world.grid.nodes[i][i].whatever = rand.Int()
//	log.Println("Updated ", i, "x", i, " to ", world.grid.nodes[i][i].whatever)
//}

func (world *World) GetGrid() Grid {
	return *(world.grid)
}

func (world *World) AddCommand(command Command) {
	world.actions.addCommand(command)
}

func (world *World) Update() {
	queue := world.actions.queue
	select {
	case cmd, ok := <-queue:
		if ok {
			world.runCommand(cmd)
		} else {
			log.Println("channel dead")
		}
	}
}

func (world *World) runCommand(command Command) {
	log.Println("Running command ", command)
	model := world.objects[command.ModelId]
	x, y := model.GetPosition()
	var didMove bool
	switch command.Direction {
	case Up:
		didMove = world.grid.move(&model, x, y-1)

	case Down:
		didMove = world.grid.move(&model, x, y+1)

	case Left:
		didMove = world.grid.move(&model, x-1, y)

	case Right:
		didMove = world.grid.move(&model, x+1, y)

	default:
		log.Fatal("world could not execute command: ", command)
	}

	if didMove {
		newX, newY := model.GetPosition()
		world.broadcastChannel <- *NewPlayerIdResponse(model.GetId(), newX, newY)
	}
}

func (world *World) StartTime() {
	log.Println("Main started game")
	world.moment.startTicking()
}

func (world *World) Stop() {
	world.moment.sendStop()
}

func (world *World) NewPerson(initX entity.GridType, initY entity.GridType) int {
	var model entity.ModelI = entity.NewPerson(initX, initY)
	world.objects[strconv.Itoa(model.GetId())] = model
	world.grid.move(&model, initX, initY)
	return model.GetId()
}

func NewWorld(size entity.GridType) (world *World) {
	world = &World{
		grid:             newGrid(size),
		actions:          Actions{queue: make(chan Command, 5)},
		objects:          make(map[string]entity.ModelI),
		broadcastChannel: make(chan PlayerIdResponse, 5),
	}
	timer := &Timer{
		world, false, 0,
	}
	world.moment = timer
	return
}
