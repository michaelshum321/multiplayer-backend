package world

type World struct {
	grid *Grid
}

func (world *World) DoSomething() {
	var numOfColumns = len(world.grid.grid)
	world.grid.grid[numOfColumns-1][0] = Node{ whatever: 69}
}

func (world *World) GetGrid() Grid{
	return *(world.grid)
}
func NewWorld(size int) *World {
	return &World{
		grid: newGrid(size),
	}
}