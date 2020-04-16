package entity

var counter = 0

type ModelS struct {
	size GridType
	x    GridType
	y    GridType
	id   int
}

type ModelI interface {
	GetSize() GridType
	GetId() int
	GetPosition() (GridType, GridType)
}


func (model *ModelS) GetSize() GridType {
	return model.size
}

func (model *ModelS) GetPosition() (GridType, GridType) {
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

func newModel(initX GridType, initY GridType, initSize GridType) ModelS {
	return ModelS{
		id:   getAndIncIdCounter(),
		x:    initX,
		y:    initY,
		size: initSize,
	}
}
