package main

import (
	"fmt"

	"claude-test/internal/creational/prototype"
)

func main() {
	fmt.Println("=== Prototype Pattern ===")

	original := prototype.NewDocumentTemplate(
		"Invoice",
		"Standard invoice template",
		map[string]string{
			"author":   "System",
			"category": "Finance",
		},
	)
	fmt.Printf("Original: %s - %s (author: %s)\n", original.Title, original.Body, original.Metadata["author"])

	cloned := original.Clone().(*prototype.DocumentTemplate)
	cloned.Title = "Custom Invoice"
	cloned.Metadata["author"] = "User"
	fmt.Printf("Clone:    %s - %s (author: %s)\n", cloned.Title, cloned.Body, cloned.Metadata["author"])
	fmt.Printf("Original: %s - %s (author: %s)\n", original.Title, original.Body, original.Metadata["author"])
}
