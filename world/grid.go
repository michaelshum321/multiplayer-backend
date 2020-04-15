package world

import "log"

type Grid struct{
	grid [][]Node
}

var initData = [][]Node{
	[]Node{
		Node{whatever: 5},
	},
	[]Node{
		Node{whatever:123},
	},
}

func newGrid(size int) (grid *Grid){
	outerGrid := make([][]Node, size, size)
	for i := range outerGrid {
		outerGrid[i]= make([]Node, size, size)
	}
	grid = &Grid{
		grid: outerGrid,
	}
	log.Println("created new grid with ", initData)
	return
}