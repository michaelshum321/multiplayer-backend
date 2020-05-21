package game

import (
	"multiplayer-backend/game/entity"
	"testing"
)

func TestStart(t *testing.T) {
	var grid = newGrid(10)
	var person entity.ModelI = entity.NewPerson(5,5)
	grid.move(&person, 5, 5)
	for _, i := range grid.nodes[4:7]{
		rowSlice := i[4:6]
		for _, j := range rowSlice {
			if j.elem != &person {
				t.Error("Node does not have reference to person")
			}
		}
	}
}

func TestMoveSimple(t *testing.T) {
	var grid = newGrid(10)
	var person entity.ModelI = entity.NewPerson(5,5)
	grid.move(&person, 5,5)
	grid.move(&person, 2,2)
	for _, i := range grid.nodes[4:7] {
		rowSlice := i[4:7]
		for _, j := range rowSlice {
			if j.elem != nil {
				t.Error("should be nil here but is ", *j.elem)
			}
		}
	}

	newX, newY := person.GetPosition()
	if newX != 2 || newY != 2 {
		t.Error("current position should be at 2x2 but is at",newX,"x",newY)
	}
	for _, i := range grid.nodes[1:4] {
		rowSlice := i[1:4]
		for _, j := range rowSlice {
			if j.elem != &person {
				t.Error("should be nil here but is ", *j.elem)
			}
		}
	}
}

func TestMoveEdges(t *testing.T) {
	grid := newGrid(10)
	var person entity.ModelI = entity.NewPerson(0,0)

	expectedX := entity.GridType(1)
	expectedY := entity.GridType(0)
	grid.move(&person, expectedX,expectedY) // move to the right

	actualX, actualY := person.GetPosition()
	if actualX != expectedX {
		t.Errorf("x should be %d but is %d", expectedX, actualX)
	}

	if actualY != expectedY {
		t.Errorf("y should be %d but is %d", expectedY, actualY)
	}
}