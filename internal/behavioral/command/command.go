/*
Package command -- Command

What is it?
Command is a behavioral design pattern that encapsulates an operation as an object, enabling
parameterization of clients with different requests, queuing of operations, and their undoing (undo) and
redoing (redo). The pattern decouples the object that invokes the operation from the object that performs it.

When to use?
  - When you need an undo/redo mechanism (e.g., text editor, graphics tools).
  - When you want to queue operations or execute them with a delay.
  - When you want to record operation history (e.g., for logging or session replay).
  - When you want to decouple the object that invokes the operation from the object that knows how to perform it.

When NOT to use?
  - When operations are simple and don't require undoing -- the pattern adds significant complexity.
  - When you don't need operation history or queuing -- direct method calls are simpler.
  - When operations are inherently irreversible (e.g., sending an email) -- an undo mechanism doesn't make sense.

Tips and pitfalls:
  - Command vs Strategy: Command stores state for undoing; Strategy is stateless.
  - ExecuteCommand clears the redo stack -- a new operation after an undo prevents redoing the undone ones.
  - DeleteCommand remembers the deleted text so it can be restored in Undo().
  - The code operates on bytes, not runes -- in production, consider using []rune instead of string.
*/
package command

import "fmt"

// Command defines the command interface with the ability to execute and undo an operation.
type Command interface {
	Execute() string
	Undo() string
}

// InsertCommand represents a command to insert text at a given position in the editor.
type InsertCommand struct {
	editor   *TextEditor
	text     string
	position int
}

// NewInsertCommand creates a new text insertion command.
func NewInsertCommand(editor *TextEditor, text string, position int) *InsertCommand {
	return &InsertCommand{editor: editor, text: text, position: position}
}

// Execute inserts text at the specified position and returns the new editor content.
func (c *InsertCommand) Execute() string {
	content := c.editor.content
	if c.position > len(content) {
		c.position = len(content)
	}
	c.editor.content = content[:c.position] + c.text + content[c.position:]
	return c.editor.content
}

// Undo reverts the text insertion and returns the restored editor content.
func (c *InsertCommand) Undo() string {
	content := c.editor.content
	c.editor.content = content[:c.position] + content[c.position+len(c.text):]
	return c.editor.content
}

// DeleteCommand represents a command to delete text from a given position in the editor.
type DeleteCommand struct {
	editor  *TextEditor
	start   int
	length  int
	deleted string
}

// NewDeleteCommand creates a new text deletion command.
func NewDeleteCommand(editor *TextEditor, start, length int) *DeleteCommand {
	return &DeleteCommand{editor: editor, start: start, length: length}
}

// Execute deletes text from the specified position, stores it, and returns the new editor content.
func (c *DeleteCommand) Execute() string {
	content := c.editor.content
	end := c.start + c.length
	if end > len(content) {
		end = len(content)
	}
	c.deleted = content[c.start:end]
	c.editor.content = content[:c.start] + content[end:]
	return c.editor.content
}

// Undo restores the previously deleted text and returns the restored editor content.
func (c *DeleteCommand) Undo() string {
	content := c.editor.content
	c.editor.content = content[:c.start] + c.deleted + content[c.start:]
	return c.editor.content
}

// TextEditor is the invoker -- it executes commands and manages operation history (undo/redo).
type TextEditor struct {
	content string
	history []Command
	undone  []Command
}

// NewTextEditor creates a new, empty text editor.
func NewTextEditor() *TextEditor {
	return &TextEditor{}
}

// ExecuteCommand executes a command, adds it to the history, and clears the redo stack.
func (e *TextEditor) ExecuteCommand(cmd Command) string {
	result := cmd.Execute()
	e.history = append(e.history, cmd)
	e.undone = nil
	return result
}

// Undo reverts the last command from the history. Returns an error if the history is empty.
func (e *TextEditor) Undo() (string, error) {
	if len(e.history) == 0 {
		return e.content, fmt.Errorf("nothing to undo")
	}
	cmd := e.history[len(e.history)-1]
	e.history = e.history[:len(e.history)-1]
	result := cmd.Undo()
	e.undone = append(e.undone, cmd)
	return result, nil
}

// Redo re-executes the last undone command. Returns an error if the redo stack is empty.
func (e *TextEditor) Redo() (string, error) {
	if len(e.undone) == 0 {
		return e.content, fmt.Errorf("nothing to redo")
	}
	cmd := e.undone[len(e.undone)-1]
	e.undone = e.undone[:len(e.undone)-1]
	result := cmd.Execute()
	e.history = append(e.history, cmd)
	return result, nil
}

// GetText returns the current text content of the editor.
func (e *TextEditor) GetText() string {
	return e.content
}
