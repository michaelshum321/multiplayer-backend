package game

import (
	"log"
	"multiplayer-backend/game/entity"
)

type GridType uint32

type Grid struct {
	nodes [][]Node
	size  GridType
}
func (grid *Grid) move(modelPtr *entity.ModelI, newX GridType, newY GridType) {
	if !grid.canMove(modelPtr,newX, newY) {
		return
	}
	nodes := grid.nodes
	model := *modelPtr
	size := model.GetSize()
	x, y := model.GetPosition()

	// TODO: optimize me :')
	// remove nodes at old position
	for i := GridType(0); i < size; i++ {
		minX, _ := getMinMax(x, size, grid.size)
		minY, maxY := getMinMax(y, size, grid.size)
		slice := nodes[minX+i][minY:maxY+1]
		for j := range slice {
			slice[j].elem = nil
		}
	}

	// set nodes at new position
	for i := GridType(0); i < size; i++ {
		minX, _ := getMinMax(newX, size, grid.size)
		minY, maxY := getMinMax(newY, size, grid.size)
		slice := nodes[minX+i][minY:maxY+1]
		for j := range slice {
			slice[j].elem = modelPtr
		}
	}
}

func getMinMax(in GridType, size GridType, maxIn GridType) (GridType, GridType) {
	offset := getOffset(in)
	radius := size/2
	var min GridType
	var max GridType
	if validBehind(in, radius, offset) {
		min = in-radius-offset
	}else{
		min = 0
	}

	if validAhead(in, maxIn, radius) {
		max = in+radius
	}else {
		max = maxIn-1
	}

	return min, max
}

/**
For Models that have odd size, offset is stored behind it.
This assumes models are only squares.
*/
func (grid *Grid) canMove(model *entity.ModelI, newX GridType, newY GridType) bool {
	radius := (*model).GetSize() / 2
	offset := getOffsetModel(model)

	return validBehind(newX, radius, offset) && validAhead(newX, grid.size, radius) &&
		validBehind(newY, radius, offset) && validAhead(newY, grid.size, radius)

}

func getOffsetModel(model *entity.ModelI) GridType{
	return getOffset((*model).GetSize())
}

func getOffset(in GridType) GridType {
	if in % 2 == 0 {
		return 0
	}
	return 1
}

//func behindOffset(in GridType, radius GridType, offset GridType) GridType {
//	if validBehind(in, radius, offset) {
//		return in-radius-offset
//	}
//	return 0
//}
func validBehind(in GridType, radius GridType, offset GridType) bool {
	return in-radius-offset >= 0 && in-radius-offset < in
}

func validAhead(in GridType, maxIn GridType, radius GridType) bool {
	return in+radius < maxIn && in+radius > in
}

func init2DNode(size GridType) (output [][]Node) {
	output = make([][]Node, size, size)
	for i := range output {
		output[i] = make([]Node, size, size)
	}
	return
}

func newGrid(size GridType) (grid *Grid) {
	grid = &Grid{
		nodes: init2DNode(size),
		size:  size,
	}
	log.Println("created new nodes with that is ", len(grid.nodes), " x ", len(grid.nodes[0]))
	return
}
