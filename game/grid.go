package game

import (
	"log"
	"multiplayer-backend/game/entity"
)

type Grid struct {
	nodes [][]Node
	size  entity.GridType
}

func (grid *Grid) PrintBoard() {
	for idx, i := range grid.nodes{
		log.Println(idx, i)
	}
}
func (grid *Grid) move(modelPtr *entity.ModelI, newX entity.GridType, newY entity.GridType) {
	if !grid.canMove(modelPtr, newX, newY) {
		log.Println("cannot move ", (*modelPtr).GetId(), " to ", newX,"x",newY)
		return
	}
	nodes := grid.nodes
	model := *modelPtr
	size := model.GetSize()
	x, y := model.GetPosition()

	log.Println("moving ", model.GetId()," from ", x,"x",y," to ", newX,"x", newY)
	grid.PrintBoard()
	// TODO: optimize me :D
	// remove nodes at old position
	minX, _ := getMinMax(x, size, grid.size)
	for i := range nodes[minX:minX+size] {
		minY, maxY := getMinMax(y, size, grid.size)
		rowSlice := nodes[entity.GridType(i)+minX][minY:maxY+1]
		for j := range rowSlice {
			rowSlice[j].elem = nil
		}
	}

	// set nodes at new position
	minX, _ = getMinMax(newX, size, grid.size)
	for  i := range nodes[minX:minX+size] {
		minY, maxY := getMinMax(newY, size, grid.size)
		rowSlice := nodes[entity.GridType(i)+minX][minY:maxY+1]
		for j := range rowSlice {
			rowSlice[j].elem = modelPtr
		}
	}

	model.SetPosition(newX, newY)
	log.Println("new grid")
	grid.PrintBoard()
}

func getMinMax(in entity.GridType, size entity.GridType, maxIn entity.GridType) (entity.GridType, entity.GridType) {
	offset := getOffset(size)
	radius := size / 2
	var min entity.GridType
	var max entity.GridType
	if validBehind(in, radius, offset) {
		min = in - radius - offset
	} else {
		min = 0
	}

	if validAhead(in, maxIn, radius) {
		max = in + radius
	} else {
		max = maxIn - 1
	}

	return min, max
}

/**
For Models that have odd size, offset is stored behind it.
This assumes models are only squares.
*/
func (grid *Grid) canMove(model *entity.ModelI, newX entity.GridType, newY entity.GridType) bool {
	radius := (*model).GetSize() / 2
	offset := getOffsetModel(model)

	return validBehind(newX, radius, offset) && validAhead(newX, grid.size, radius) &&
		validBehind(newY, radius, offset) && validAhead(newY, grid.size, radius)
}

func getOffsetModel(model *entity.ModelI) entity.GridType {
	return getOffset((*model).GetSize())
}

func getOffset(in entity.GridType) entity.GridType {
	if in%2 == 0 {
		return 1
	}
	return 0
}

//func behindOffset(in GridType, radius GridType, offset GridType) GridType {
//	if validBehind(in, radius, offset) {
//		return in-radius-offset
//	}
//	return 0
//}
func validBehind(in entity.GridType, radius entity.GridType, offset entity.GridType) bool {
	return in-radius-offset >= 0 && in-radius-offset <= in
}

func validAhead(in entity.GridType, maxIn entity.GridType, radius entity.GridType) bool {
	return in+radius < maxIn && in+radius >= in
}

func init2DNode(size entity.GridType) (output [][]Node) {
	output = make([][]Node, size, size)
	for i := range output {
		output[i] = make([]Node, size, size)
	}
	return
}

func newGrid(size entity.GridType) (grid *Grid) {
	grid = &Grid{
		nodes: init2DNode(size),
		size:  size,
	}
	log.Println("created new nodes with that is ", len(grid.nodes), " x ", len(grid.nodes[0]))
	return
}
