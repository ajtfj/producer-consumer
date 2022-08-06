package eventbuffer

import (
	"runtime"
	"sync"
)

type EventBuffer struct {
	mu     *sync.Mutex
	cap    int
	buffer []int
}

func NewEventBuffer(cap int) *EventBuffer {
	eventBuffer := EventBuffer{
		mu:  &sync.Mutex{},
		cap: cap,
	}

	return &eventBuffer
}

func (eb *EventBuffer) Consume() int {
	eb.mu.Lock()
	for len(eb.buffer) == 0 {
		eb.mu.Unlock()
		runtime.Gosched()
		eb.mu.Lock()
	}
	defer eb.mu.Unlock()

	event := eb.buffer[0]
	eb.buffer = eb.buffer[1:]
	return event
}

func (eb *EventBuffer) Produce(event int) {
	eb.mu.Lock()
	for len(eb.buffer) == eb.cap {
		eb.mu.Unlock()
		runtime.Gosched()
		eb.mu.Lock()
	}
	defer eb.mu.Unlock()

	eb.buffer = append(eb.buffer, event)
}
