package main

import (
	"fmt"

	"claude-test/internal/behavioral/observer"
)

type logger struct {
	name string
}

func (l *logger) OnEvent(e observer.Event) {
	fmt.Printf("  [%s] received %s: %s\n", l.name, e.Type, e.Data)
}

func main() {
	fmt.Println("=== Observer Pattern ===")

	bus := observer.NewEventBus()

	emailNotifier := &logger{name: "EmailNotifier"}
	auditLog := &logger{name: "AuditLog"}
	analytics := &logger{name: "Analytics"}

	bus.Subscribe("user.created", emailNotifier)
	bus.Subscribe("user.created", auditLog)
	bus.Subscribe("order.placed", auditLog)
	bus.Subscribe("order.placed", analytics)

	fmt.Println("Publishing user.created:")
	bus.Publish(observer.Event{Type: "user.created", Data: "alice@example.com"})

	fmt.Println("Publishing order.placed:")
	bus.Publish(observer.Event{Type: "order.placed", Data: "order-123"})
}
