package inmemory

import (
	"context"

	"github.com/ArthurQR98/challenge_fiber/kit/event"
)

type EventBus struct {
	events []event.Event
}

func NewEventBus() *EventBus {
	return &EventBus{}
}

func (b *EventBus) Publish(_ context.Context, events []event.Event) error {
	b.events = append(b.events, events...)
	return nil
}
