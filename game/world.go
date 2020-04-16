package game

import (
	"log"
	"multiplayer-backend/game/entity"
	"strconv"
)

type World struct {
	grid   *Grid
	moment *Timer
	Actions
	objects map[string]entity.ModelI
}

//func (world *World) doSomething() {
//	i := GridType(rand.Uint32()) % world.grid.size
//	world.grid.nodes[i][i].whatever = rand.Int()
//	log.Println("Updated ", i, "x", i, " to ", world.grid.nodes[i][i].whatever)
//}

func (world *World) GetGrid() Grid {
	return *(world.grid)
}

func (world *World) Update() {
	queue := world.Actions.GetQueue()
	select {
	case cmd, ok := <-queue:
		if ok {
			log.Println("channel has something")
			world.runCommand(cmd)
		} else {
			log.Println("channel dead")
		}
	default:
		log.Println("no value")
	}
}

func (world *World) runCommand(command Command) {
	model := world.objects[command.modelId]
	world.grid
}

func (world *World) StartTime() {
	world.moment.startTicking()
}

func (world *World) Stop() {
	world.moment.sendStop()
}

func (world *World) NewPerson(initX GridType, initY GridType) {
	var model entity.ModelI = entity.NewPerson(initX, initY)
	world.objects[strconv.Itoa(model.GetId())] = model
}

func NewWorld(size GridType) (world *World) {
	world = &World{
		grid:    newGrid(size),
		Actions: Actions{queue: make(chan Command, 100)},
		objects: make(map[string]entity.ModelI),
	}
	timer := &Timer{
		world, false, 0,
	}
	world.moment = timer
	return
}
