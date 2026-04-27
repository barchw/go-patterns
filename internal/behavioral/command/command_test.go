package command

import "testing"

func TestTextEditor_ExecuteInsert(t *testing.T) {
	editor := NewTextEditor()

	result := editor.ExecuteCommand(NewInsertCommand(editor, "Hello", 0))
	if result != "Hello" {
		t.Errorf("got %q, want %q", result, "Hello")
	}
	if editor.GetText() != "Hello" {
		t.Errorf("GetText() = %q, want %q", editor.GetText(), "Hello")
	}
}

func TestTextEditor_ExecuteDelete(t *testing.T) {
	editor := NewTextEditor()
	editor.ExecuteCommand(NewInsertCommand(editor, "Hello World", 0))

	result := editor.ExecuteCommand(NewDeleteCommand(editor, 5, 6))
	if result != "Hello" {
		t.Errorf("got %q, want %q", result, "Hello")
	}
}

func TestTextEditor_Undo(t *testing.T) {
	editor := NewTextEditor()
	editor.ExecuteCommand(NewInsertCommand(editor, "Hello", 0))
	editor.ExecuteCommand(NewInsertCommand(editor, " World", 5))

	result, err := editor.Undo()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Hello" {
		t.Errorf("after first undo: got %q, want %q", result, "Hello")
	}

	result, err = editor.Undo()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("after second undo: got %q, want %q", result, "")
	}
}

func TestTextEditor_Redo(t *testing.T) {
	editor := NewTextEditor()
	editor.ExecuteCommand(NewInsertCommand(editor, "Hello", 0))
	editor.ExecuteCommand(NewInsertCommand(editor, " World", 5))
	editor.Undo()
	editor.Undo()

	result, err := editor.Redo()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Hello" {
		t.Errorf("after first redo: got %q, want %q", result, "Hello")
	}

	result, err = editor.Redo()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "Hello World" {
		t.Errorf("after second redo: got %q, want %q", result, "Hello World")
	}
}

func TestTextEditor_RedoClearedOnNewCommand(t *testing.T) {
	editor := NewTextEditor()
	editor.ExecuteCommand(NewInsertCommand(editor, "Hello", 0))
	editor.Undo()
	editor.ExecuteCommand(NewInsertCommand(editor, "Bye", 0))

	_, err := editor.Redo()
	if err == nil {
		t.Error("expected error on redo after new command, got nil")
	}
}

func TestTextEditor_UndoEmptyHistory(t *testing.T) {
	editor := NewTextEditor()
	_, err := editor.Undo()
	if err == nil {
		t.Error("expected error on undo with empty history, got nil")
	}
}

func TestTextEditor_RedoEmptyHistory(t *testing.T) {
	editor := NewTextEditor()
	_, err := editor.Redo()
	if err == nil {
		t.Error("expected error on redo with empty history, got nil")
	}
}

func TestTextEditor_MultipleCommandSequence(t *testing.T) {
	editor := NewTextEditor()

	editor.ExecuteCommand(NewInsertCommand(editor, "Hello World", 0))
	if editor.GetText() != "Hello World" {
		t.Fatalf("after insert: got %q", editor.GetText())
	}

	editor.ExecuteCommand(NewDeleteCommand(editor, 5, 6))
	if editor.GetText() != "Hello" {
		t.Fatalf("after delete: got %q", editor.GetText())
	}

	editor.ExecuteCommand(NewInsertCommand(editor, " Go", 5))
	if editor.GetText() != "Hello Go" {
		t.Fatalf("after second insert: got %q", editor.GetText())
	}

	editor.Undo()
	if editor.GetText() != "Hello" {
		t.Fatalf("after undo insert: got %q", editor.GetText())
	}

	editor.Undo()
	if editor.GetText() != "Hello World" {
		t.Fatalf("after undo delete: got %q", editor.GetText())
	}

	editor.Redo()
	if editor.GetText() != "Hello" {
		t.Fatalf("after redo delete: got %q", editor.GetText())
	}
}
