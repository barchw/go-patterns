package main

import (
	"fmt"

	"claude-test/internal/structural/facade"
)

func main() {
	fmt.Println("=== Facade Pattern ===")

	converter := facade.NewMediaConverter()
	result := converter.Convert("vacation.avi", "mp4")
	fmt.Println("Conversion result:", result)
}
