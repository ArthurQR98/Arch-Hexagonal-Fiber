package inmemory

import (
	"context"

	"github.com/ArthurQR98/challenge_fiber/kit/event"
)

type EventBus struct {
	handlers map[event.Type][]event.Handler
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[event.Type][]event.Handler),
	}
}

func (b *EventBus) Publish(ctx context.Context, events []event.Event) error {
	for _, evt := range events {
		handlers, ok := b.handlers[evt.Type()]
		if !ok {
			return nil
		}

		for _, handler := range handlers {
			handler.Handle(ctx, evt)
		}
	}
	return nil
}

func (b *EventBus) Subscribe(evtType event.Type, handler event.Handler) {
	subscribersForType, ok := b.handlers[evtType]
	if !ok {
		b.handlers[evtType] = []event.Handler{handler}
	}

	subscribersForType = append(subscribersForType, handler)
}
