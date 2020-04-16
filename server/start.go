package main

import (
	"log"
	"multiplayer-backend/game"
	"time"
)

func main() {
	var world = game.NewWorld(5)
	go world.StartTime()
	log.Println("Main started game")

	time.Sleep(time.Minute/6)
	world.Stop()
	time.Sleep(time.Second)
	log.Println("exiting Main now")
}
