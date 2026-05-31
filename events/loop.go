package events

import (
	"tui-engine/backend"
)

type EventLoop struct {
	backend backend.Backend
}

func NewEventLoop(backend backend.Backend) *EventLoop {
	return &EventLoop{backend: backend}
}

func (l *EventLoop) Start() <-chan Event {
	channel := make(chan Event)

	go watchKeyEvents(channel)
	go watchResizeEvents(channel)

	return channel
}
