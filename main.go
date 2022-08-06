package main

import (
	"fmt"
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

func consumer(eb *eventbuffer.EventBuffer, ch *chan int) {
	event := eb.Consume()
	fmt.Println(event)
	<-*ch
}

func main() {
	eventBuffer := eventbuffer.NewEventBuffer(CAP)

	producerWg := sync.WaitGroup{}
	producerWg.Add(1)
	go producer(eventBuffer, &producerWg)

	chConsumer := make(chan int, MAX_CONSUMERS)
	for i := 0; i < MAX_EVENTS; i++ {
		chConsumer <- i
		go consumer(eventBuffer, &chConsumer)
	}

	producerWg.Wait()
}
