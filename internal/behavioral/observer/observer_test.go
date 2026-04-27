package observer

import "testing"

type mockSubscriber struct {
	received []Event
}

func (m *mockSubscriber) OnEvent(e Event) {
	m.received = append(m.received, e)
}

func TestEventBus_AllSubscribersNotified(t *testing.T) {
	bus := NewEventBus()
	sub1 := &mockSubscriber{}
	sub2 := &mockSubscriber{}

	bus.Subscribe("user.created", sub1)
	bus.Subscribe("user.created", sub2)

	event := Event{Type: "user.created", Data: "alice"}
	bus.Publish(event)

	if len(sub1.received) != 1 {
		t.Errorf("sub1 received %d events, want 1", len(sub1.received))
	}
	if len(sub2.received) != 1 {
		t.Errorf("sub2 received %d events, want 1", len(sub2.received))
	}
	if sub1.received[0].Data != "alice" {
		t.Errorf("sub1 got data %q, want %q", sub1.received[0].Data, "alice")
	}
	if sub2.received[0].Data != "alice" {
		t.Errorf("sub2 got data %q, want %q", sub2.received[0].Data, "alice")
	}
}

func TestEventBus_OnlyMatchingTypeSubscribersCalled(t *testing.T) {
	bus := NewEventBus()
	userSub := &mockSubscriber{}
	orderSub := &mockSubscriber{}

	bus.Subscribe("user.created", userSub)
	bus.Subscribe("order.placed", orderSub)

	bus.Publish(Event{Type: "user.created", Data: "bob"})

	if len(userSub.received) != 1 {
		t.Errorf("userSub received %d events, want 1", len(userSub.received))
	}
	if len(orderSub.received) != 0 {
		t.Errorf("orderSub received %d events, want 0", len(orderSub.received))
	}
}

func TestEventBus_MultipleEventTypes(t *testing.T) {
	bus := NewEventBus()
	sub := &mockSubscriber{}

	bus.Subscribe("a", sub)
	bus.Subscribe("b", sub)

	bus.Publish(Event{Type: "a", Data: "1"})
	bus.Publish(Event{Type: "b", Data: "2"})
	bus.Publish(Event{Type: "c", Data: "3"})

	if len(sub.received) != 2 {
		t.Errorf("sub received %d events, want 2", len(sub.received))
	}
}

func TestEventBus_NoSubscribers(t *testing.T) {
	bus := NewEventBus()
	bus.Publish(Event{Type: "unknown", Data: "data"})
}
