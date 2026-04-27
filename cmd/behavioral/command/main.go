package main

import (
	"fmt"

	"claude-test/internal/behavioral/command"
)

func main() {
	fmt.Println("=== Command Pattern ===")

	editor := command.NewTextEditor()

	fmt.Printf("Insert 'Hello':      %q\n", editor.ExecuteCommand(command.NewInsertCommand(editor, "Hello", 0)))
	fmt.Printf("Insert ' World':     %q\n", editor.ExecuteCommand(command.NewInsertCommand(editor, " World", 5)))
	fmt.Printf("Delete ' World':     %q\n", editor.ExecuteCommand(command.NewDeleteCommand(editor, 5, 6)))

	result, _ := editor.Undo()
	fmt.Printf("Undo (delete):       %q\n", result)

	result, _ = editor.Undo()
	fmt.Printf("Undo (insert):       %q\n", result)

	result, _ = editor.Redo()
	fmt.Printf("Redo (insert):       %q\n", result)
}
