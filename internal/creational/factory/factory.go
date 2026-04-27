/*
Package factory implements the Factory Method pattern.

What is it?
Factory Method defines an interface for creating objects but lets the factory
function decide which concrete implementation to create. Client code operates
solely on the interface, without knowing the implementation details.

When to use?
  - When you don't know in advance what type of object will be needed -- the decision is made at runtime.
  - When you want to separate the object-creation code from the code that uses the objects.
  - When a switch/if-else appears that decides which type of object to create.
  - When different implementations have different dependencies or initialization logic.

When NOT to use?
  - When there is only one implementation type and simple construction is sufficient.
  - When the factory artificially complicates the code and the number of variants is minimal.
  - When the object requires no initialization logic.

Tips and pitfalls:
  - Unexported structs force the use of the factory -- this is natural encapsulation in Go.
  - Always return an error from the factory function to handle unknown types.
  - To add a new variant, you only need a new constant, a struct, and a case in the switch -- client code requires no changes.
  - In larger projects, consider a registry (map) instead of a growing switch.
*/
package factory

import "fmt"

// Notification defines the notification interface -- a contract that all notification types must satisfy.
type Notification interface {
	Send(to, message string) string
}

// NotificationType represents an enum type that controls the notification factory.
type NotificationType int

const (
	// EmailNotification represents an email notification.
	EmailNotification NotificationType = iota
	// SMSNotification represents an SMS notification.
	SMSNotification
	// PushNotification represents a push notification.
	PushNotification
)

type emailNotification struct{}

func (e *emailNotification) Send(to, message string) string {
	return fmt.Sprintf("Email to %s: %s", to, message)
}

type smsNotification struct{}

func (s *smsNotification) Send(to, message string) string {
	return fmt.Sprintf("SMS to %s: %s", to, message)
}

type pushNotification struct{}

func (p *pushNotification) Send(to, message string) string {
	return fmt.Sprintf("Push to %s: %s", to, message)
}

// NewNotification creates a notification of the appropriate type based on the given NotificationType.
func NewNotification(nt NotificationType) (Notification, error) {
	switch nt {
	case EmailNotification:
		return &emailNotification{}, nil
	case SMSNotification:
		return &smsNotification{}, nil
	case PushNotification:
		return &pushNotification{}, nil
	default:
		return nil, fmt.Errorf("unknown notification type: %d", nt)
	}
}
