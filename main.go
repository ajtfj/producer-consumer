package main

import (
	"sync"

	"github.com/ajtfj/producer-consumer/eventbuffer"
)

const (
	MAX_CONSUMERS = 100
	MAX_EVENTS    = 1000
	CAP           = 1
)

func producer(eb *eventbuffer.EventBuffer, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < MAX_EVENTS; i++ {
		eb.Produce(i)
	}
}

func main() {
	eventBuffer := eventbuffer.NewEventBuffer(CAP)

	producerWg := sync.WaitGroup{}
	producerWg.Add(1)
	go producer(eventBuffer, &producerWg)

	producerWg.Wait()
}
