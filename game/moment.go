package game

import (
	"log"
	"time"
)

type Timer struct {
	world       *World
	stop        bool
	timeElapsed uint64
}

func (timer *Timer) startTicking() {
	for range time.Tick(time.Second) {
		timer.tick()
		if timer.stop {
			log.Println("received stop signal @ timeElapsed", timer.timeElapsed)
			return
		}
		timer.world.Update()
	}
}

func (timer *Timer) tick() {
	timer.timeElapsed++
	log.Println("ticked @", timer.timeElapsed)
}

func (timer *Timer) sendStop() {
	timer.stop = true
	log.Println("stop signal sent")
}
