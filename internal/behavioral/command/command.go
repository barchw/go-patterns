package command

import "fmt"

type Command interface {
	Execute() string
	Undo() string
}

type InsertCommand struct {
	editor   *TextEditor
	text     string
	position int
}

func NewInsertCommand(editor *TextEditor, text string, position int) *InsertCommand {
	return &InsertCommand{editor: editor, text: text, position: position}
}

func (c *InsertCommand) Execute() string {
	content := c.editor.content
	if c.position > len(content) {
		c.position = len(content)
	}
	c.editor.content = content[:c.position] + c.text + content[c.position:]
	return c.editor.content
}

func (c *InsertCommand) Undo() string {
	content := c.editor.content
	c.editor.content = content[:c.position] + content[c.position+len(c.text):]
	return c.editor.content
}

type DeleteCommand struct {
	editor  *TextEditor
	start   int
	length  int
	deleted string
}

func NewDeleteCommand(editor *TextEditor, start, length int) *DeleteCommand {
	return &DeleteCommand{editor: editor, start: start, length: length}
}

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

func (c *DeleteCommand) Undo() string {
	content := c.editor.content
	c.editor.content = content[:c.start] + c.deleted + content[c.start:]
	return c.editor.content
}

type TextEditor struct {
	content string
	history []Command
	undone  []Command
}

func NewTextEditor() *TextEditor {
	return &TextEditor{}
}

func (e *TextEditor) ExecuteCommand(cmd Command) string {
	result := cmd.Execute()
	e.history = append(e.history, cmd)
	e.undone = nil
	return result
}

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

func (e *TextEditor) GetText() string {
	return e.content
}
