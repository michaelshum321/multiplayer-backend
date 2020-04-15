package world

import "log"

type Grid struct{
	nodes [][]Node
	size uint32
}

func init2DNode(size uint32) (output [][]Node) {
	output = make([][]Node, size, size)
	for i := range output {
		output[i]= make([]Node, size, size)
	}
	return
}
func newGrid(size uint32) (grid *Grid){
	grid = &Grid{
		nodes: init2DNode(size),
		size: size,
	}
	log.Println("created new nodes with that is ", len(grid.nodes), " x ", len(grid.nodes[0]))
	return
}