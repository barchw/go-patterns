package main

import (
	"fmt"

	"claude-test/internal/creational/singleton"
)

func main() {
	fmt.Println("=== Singleton Pattern ===")

	cfg := singleton.GetInstance()
	cfg.Set("app.name", "MyApp")
	cfg.Set("app.version", "1.0.0")

	cfg2 := singleton.GetInstance()
	fmt.Printf("app.name = %s\n", cfg2.Get("app.name"))
	fmt.Printf("app.version = %s\n", cfg2.Get("app.version"))
	fmt.Printf("same instance: %v\n", cfg == cfg2)
}
