package main

import (
	"math/rand"
	"time"

	"github.com/ajtfj/producer-consumer/eventbuffer"
)

const (
	MAX_CONSUMERS = 100
	MAX_EVENTS    = 1000
	CAP           = 1
)

func generateEvent() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func producer(eb *eventbuffer.EventBuffer) {
	for i := 0; i < MAX_EVENTS; i++ {
		event := generateEvent()
		eb.Produce(event)
	}
}

func main() {
	eventBuffer := eventbuffer.NewEventBuffer(CAP)
	go producer(eventBuffer)
}
