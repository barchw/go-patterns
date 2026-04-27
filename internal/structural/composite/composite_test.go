package composite

import (
	"strings"
	"testing"
)

func TestFileComponent(t *testing.T) {
	f := NewFile("readme.txt", 100)

	if f.GetName() != "readme.txt" {
		t.Fatalf("got name %q, want %q", f.GetName(), "readme.txt")
	}
	if f.GetSize() != 100 {
		t.Fatalf("got size %d, want %d", f.GetSize(), 100)
	}
	if got := f.Display(""); got != "readme.txt (100 bytes)" {
		t.Fatalf("got display %q, want %q", got, "readme.txt (100 bytes)")
	}
}

func TestDirectorySizeAggregation(t *testing.T) {
	root := NewDirectory("root")
	root.Add(NewFile("a.txt", 50))
	root.Add(NewFile("b.txt", 150))

	sub := NewDirectory("sub")
	sub.Add(NewFile("c.txt", 300))
	root.Add(sub)

	if got := root.GetSize(); got != 500 {
		t.Fatalf("got size %d, want %d", got, 500)
	}
}

func TestDirectoryDisplay(t *testing.T) {
	root := NewDirectory("project")
	root.Add(NewFile("main.go", 200))

	src := NewDirectory("src")
	src.Add(NewFile("app.go", 400))
	root.Add(src)

	output := root.Display("")

	expectedParts := []string{
		"project/",
		"  main.go (200 bytes)",
		"  src/",
		"    app.go (400 bytes)",
	}

	for _, part := range expectedParts {
		if !strings.Contains(output, part) {
			t.Fatalf("display output missing %q.\nFull output:\n%s", part, output)
		}
	}
}

func TestEmptyDirectory(t *testing.T) {
	d := NewDirectory("empty")
	if d.GetSize() != 0 {
		t.Fatalf("empty dir size should be 0, got %d", d.GetSize())
	}
	if got := d.Display(""); got != "empty/" {
		t.Fatalf("got display %q, want %q", got, "empty/")
	}
}

func TestComponentInterface(t *testing.T) {
	var c Component
	c = NewFile("f.txt", 10)
	_ = c.GetName()
	c = NewDirectory("d")
	_ = c.GetName()
}
