package main

import (
	"fmt"
	"log"

	"claude-test/internal/creational/factory"
)

func main() {
	fmt.Println("=== Factory Method Pattern ===")

	types := []struct {
		name  string
		ntype factory.NotificationType
	}{
		{"Email", factory.EmailNotification},
		{"SMS", factory.SMSNotification},
		{"Push", factory.PushNotification},
	}

	for _, t := range types {
		n, err := factory.NewNotification(t.ntype)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(n.Send("user@example.com", "Hello from "+t.name))
	}
}
