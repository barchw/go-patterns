package factory

import "fmt"

type Notification interface {
	Send(to, message string) string
}

type NotificationType int

const (
	EmailNotification NotificationType = iota
	SMSNotification
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
