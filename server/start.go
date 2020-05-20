package main

import (
	"log"
	"multiplayer-backend/game"
	"time"
)

func main() {
	var worldPtr = game.NewWorld(10)
	var connHandler = game.NewGameConnectionHandler(worldPtr)
	var tcpServer = NewWSServer(connHandler)
	tcpServer.Start()
	// go worldPtr.AddCommand(game.Command{ModelId: "0", Dir: game.Right})
	// go worldPtr.AddCommand(game.Command{ModelId: "0", Dir: game.Left})
	time.Sleep(time.Minute*5)
	worldPtr.Stop()
	time.Sleep(time.Second)
	log.Println("exiting Main now")
}
