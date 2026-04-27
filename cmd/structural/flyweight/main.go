package main

import (
	"fmt"

	"claude-test/internal/structural/flyweight"
)

func main() {
	fmt.Println("=== Flyweight Pattern ===")

	factory := flyweight.NewGlyphFactory()
	renderer := flyweight.NewCharacterRenderer(factory)

	text := "HELLO"
	for i, ch := range text {
		fmt.Println(renderer.Render(ch, "Arial", i*12, 0))
	}

	fmt.Printf("\nRendered %d characters using %d unique glyphs\n",
		renderer.CharacterCount(), factory.GlyphCount())
}
