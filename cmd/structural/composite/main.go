package main

import (
	"fmt"

	"claude-test/internal/structural/composite"
)

func main() {
	fmt.Println("=== Composite Pattern ===")

	root := composite.NewDirectory("project")
	root.Add(composite.NewFile("go.mod", 50))
	root.Add(composite.NewFile("main.go", 200))

	src := composite.NewDirectory("src")
	src.Add(composite.NewFile("app.go", 400))
	src.Add(composite.NewFile("util.go", 150))
	root.Add(src)

	tests := composite.NewDirectory("tests")
	tests.Add(composite.NewFile("app_test.go", 300))
	root.Add(tests)

	fmt.Println(root.Display(""))
	fmt.Printf("\nTotal size: %d bytes\n", root.GetSize())
}
