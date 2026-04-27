package main

import (
	"fmt"

	af "claude-test/internal/creational/abstract_factory"
)

func renderUI(factory af.UIFactory) {
	btn := factory.CreateButton("Submit")
	cb := factory.CreateCheckbox("Remember me")
	fmt.Println(btn.Render())
	fmt.Println(cb.Render())
}

func main() {
	fmt.Println("=== Abstract Factory Pattern ===")

	fmt.Println("Light theme:")
	renderUI(&af.LightThemeFactory{})

	fmt.Println("Dark theme:")
	renderUI(&af.DarkThemeFactory{})
}
