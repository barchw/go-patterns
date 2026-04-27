package observer

import "sync"

type Event struct {
	Type string
	Data string
}

type Subscriber interface {
	OnEvent(Event)
}

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]Subscriber
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Subscriber),
	}
}

func (eb *EventBus) Subscribe(eventType string, subscriber Subscriber) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], subscriber)
}

func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, sub := range eb.subscribers[event.Type] {
		sub.OnEvent(event)
	}
}
