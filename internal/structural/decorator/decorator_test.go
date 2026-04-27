package decorator

import (
	"strings"
	"testing"
)

func TestFileDataSource(t *testing.T) {
	ds := NewFileDataSource("test.txt")
	ds.WriteData("hello world")
	if got := ds.ReadData(); got != "hello world" {
		t.Fatalf("got %q, want %q", got, "hello world")
	}
}

func TestCompressionDecorator(t *testing.T) {
	base := NewFileDataSource("test.txt")
	var ds DataSource = NewCompressionDecorator(base)

	result := ds.WriteData("hello world")
	if !strings.Contains(result, "[Compressed]") {
		t.Fatalf("expected [Compressed] in write result: %s", result)
	}

	if got := ds.ReadData(); got != "hello world" {
		t.Fatalf("got %q, want %q", got, "hello world")
	}
}

func TestEncryptionDecorator(t *testing.T) {
	base := NewFileDataSource("test.txt")
	var ds DataSource = NewEncryptionDecorator(base)

	result := ds.WriteData("secret data")
	if !strings.Contains(result, "[Encrypted]") {
		t.Fatalf("expected [Encrypted] in write result: %s", result)
	}

	raw := base.ReadData()
	if raw == "secret data" {
		t.Fatal("data should be encrypted in the underlying source")
	}

	if got := ds.ReadData(); got != "secret data" {
		t.Fatalf("got %q, want %q", got, "secret data")
	}
}

func TestStackedDecorators(t *testing.T) {
	base := NewFileDataSource("test.txt")
	compressed := NewCompressionDecorator(base)
	var ds DataSource = NewEncryptionDecorator(compressed)

	result := ds.WriteData("important message")
	if !strings.Contains(result, "[Encrypted]") {
		t.Fatalf("expected [Encrypted] in write result: %s", result)
	}
	if !strings.Contains(result, "[Compressed]") {
		t.Fatalf("expected [Compressed] in write result: %s", result)
	}

	if got := ds.ReadData(); got != "important message" {
		t.Fatalf("got %q, want %q", got, "important message")
	}
}

func TestDecoratorPreservesInterface(t *testing.T) {
	base := NewFileDataSource("f.txt")
	var ds DataSource = base
	ds = NewCompressionDecorator(ds)
	ds = NewEncryptionDecorator(ds)

	_ = ds.WriteData("test")
	_ = ds.ReadData()
}
