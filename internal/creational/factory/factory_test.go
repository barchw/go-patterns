package factory

import "testing"

func TestNewNotification(t *testing.T) {
	tests := []struct {
		name     string
		ntype    NotificationType
		to       string
		message  string
		expected string
	}{
		{
			name:     "email notification",
			ntype:    EmailNotification,
			to:       "user@example.com",
			message:  "Hello",
			expected: "Email to user@example.com: Hello",
		},
		{
			name:     "sms notification",
			ntype:    SMSNotification,
			to:       "+1234567890",
			message:  "Hi there",
			expected: "SMS to +1234567890: Hi there",
		},
		{
			name:     "push notification",
			ntype:    PushNotification,
			to:       "device-token-abc",
			message:  "New update",
			expected: "Push to device-token-abc: New update",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n, err := NewNotification(tt.ntype)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			got := n.Send(tt.to, tt.message)
			if got != tt.expected {
				t.Errorf("got %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestNewNotification_Unknown(t *testing.T) {
	_, err := NewNotification(NotificationType(99))
	if err == nil {
		t.Fatal("expected error for unknown notification type")
	}
}
