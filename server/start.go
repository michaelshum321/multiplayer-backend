package main

import (
	"log"
	"multiplayer-backend/world"
	"time"
)

func main() {
	var world = world.NewWorld(5)
	go world.StartTime()
	log.Println("Main started world")
	time.Sleep(time.Minute/6)
	world.Stop()
	time.Sleep(time.Second)
	log.Println("exiting Main now")
}
