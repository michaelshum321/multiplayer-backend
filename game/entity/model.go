package entity

import (
	"multiplayer-backend/game"
)

var counter = 0

type ModelS struct {
	size game.GridType
	x    game.GridType
	y    game.GridType
	id   int
}

type ModelI interface {
	GetSize() game.GridType
	GetId() int
	GetPosition() (game.GridType, game.GridType)
}


func (model *ModelS) GetSize() game.GridType {
	return model.size
}

func (model *ModelS) GetPosition() (game.GridType, game.GridType) {
	return model.x, model.y
}

func (model *ModelS) GetId() int {
	return model.id
}

func getAndIncIdCounter() int {
	id := counter
	counter++
	return id
}

func newModel(initX game.GridType, initY game.GridType, initSize game.GridType) ModelS {
	return ModelS{
		id:   getAndIncIdCounter(),
		x:    initX,
		y:    initY,
		size: initSize,
	}
}
