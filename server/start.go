package main

import(
	"log"
	"multiplayer-backend/world")

func main() {
	var world = world.NewWorld(5)
	log.Println("old world: ", world, world.GetGrid())
	world.DoSomething()
	log.Println("new world: ", world, world.GetGrid())
}
