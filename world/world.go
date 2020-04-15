package world

import (
	"log"
	"math/rand"
)

type World struct {
	grid *Grid
	moment *Timer
}

func (world *World) doSomething() {
	i := rand.Uint32() % world.grid.size
	world.grid.nodes[i][i].whatever = rand.Int()
	log.Println("Updated ",i,"x",i, " to ", world.grid.nodes[i][i].whatever)
}

func (world *World) GetGrid() Grid {
	return *(world.grid)
}

func (world *World) Update() {
	world.doSomething()
}

func (world *World) StartTime() {
	world.moment.startTicking()
}

func (world *World) Stop() {
	world.moment.sendStop()
}

func NewWorld(size uint32) (world *World) {
	world = &World{
		grid: newGrid(size),
	}
	timer := &Timer{
		world, false, 0,
	}
	world.moment = timer
	return
}
