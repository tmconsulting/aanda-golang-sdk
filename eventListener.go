package aandaSdk

import (
	"context"
	"sync"
)

const (
	BeforeRequestSend    EventType = 1
	AfterResponseReceive EventType = 2
)

type EventType int

type EventHandler func(ctx context.Context, methodName, mimeType string, data []byte)

type EventListener struct {
	listeners map[EventType][]EventHandler
}

func (o *EventListener) Init() {
	o.listeners = make(map[EventType][]EventHandler)
}

func (o *EventListener) RegisterEventHandler(et EventType, handler EventHandler) *EventListener {
	o.listeners[et] = append(o.listeners[et], handler)

	return o
}

func (o *EventListener) raiseEvent(et EventType, ctx context.Context, methodName, mimeType string, data []byte) {
	var wait sync.WaitGroup

	for _, handler := range o.listeners[et] {
		wait.Add(1)
		go func() {
			defer wait.Done()

			handler(ctx, methodName, mimeType, data)
		}()
	}

	wait.Wait()
}
