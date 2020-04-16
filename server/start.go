package main

import (
	"log"
	"multiplayer-backend/game"
	"time"
)

func main() {
	var world = game.NewWorld(10)
	go world.StartTime()
	log.Println("Main started game")
	world.NewPerson(2,2) // id 0
	go world.AddCommand(game.Command{ModelId: "0", Dir: game.Right})
	go world.AddCommand(game.Command{ModelId: "0", Dir: game.Left})
	time.Sleep(time.Minute/6)
	world.Stop()
	time.Sleep(time.Second)
	log.Println("exiting Main now")
}
