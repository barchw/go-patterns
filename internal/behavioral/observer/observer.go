/*
Package observer -- Observer

What is it?
Observer is a behavioral design pattern that defines a subscription mechanism allowing
multiple objects to listen for and react to events occurring in another object. This implementation
uses the Event Bus form, which enables subscription by event type.

When to use?
  - When a change in one object's state should automatically notify other objects.
  - When you want to decouple event senders from their receivers (loose coupling).
  - When you're building event-driven systems, e.g., notifications, logging, analytics.
  - When multiple independent modules need to react to the same event in different ways.

When NOT to use?
  - When the order of subscriber notification matters -- the pattern does not guarantee order.
  - When communication is bidirectional or requires a response -- Observer is unidirectional.
  - When you have a very small number of fixed receivers -- direct method calls are simpler.

Tips and pitfalls:
  - sync.RWMutex allows concurrent reads (publishes), blocking only writes (subscriptions).
  - No Unsubscribe mechanism -- in production, consider adding one, e.g., by returning a cancel function.
  - Publish calls subscribers synchronously -- a slow subscriber will block the entire publish.
  - Observer vs Mediator: Mediator centralizes communication; Observer is a sender-receiver relationship.
*/
package observer

import "sync"

// Event represents an event with a type (filtering key) and data.
type Event struct {
	Type string
	Data string
}

// Subscriber is an interface that every event receiver must implement.
type Subscriber interface {
	OnEvent(Event)
}

// EventBus is a central event bus that holds subscribers grouped by event type.
type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]Subscriber
}

// NewEventBus creates a new event bus with an initialized subscriber map.
func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]Subscriber),
	}
}

// Subscribe registers a subscriber for a given event type. The operation is thread-safe.
func (eb *EventBus) Subscribe(eventType string, subscriber Subscriber) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	eb.subscribers[eventType] = append(eb.subscribers[eventType], subscriber)
}

// Publish publishes an event to all subscribers registered for its type.
func (eb *EventBus) Publish(event Event) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, sub := range eb.subscribers[event.Type] {
		sub.OnEvent(event)
	}
}
