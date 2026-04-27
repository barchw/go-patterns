/*
Package composite implements the Composite design pattern.

What is it?
Composite lets you treat individual objects and groups of objects in a uniform way.
It builds tree structures where both leaves and internal nodes implement the same
interface. In this package, the pattern models a file system -- File is a leaf,
and Directory is a container that holds files and other directories.

When to use?
  - When you need to represent tree hierarchies (file systems, DOM trees, menus).
  - When client code should treat individual elements and groups identically.
  - When operations on the structure (e.g., calculating size) should be performed recursively.
  - When new element types can be added without changing client code.

When NOT to use?
  - When the structure is flat and has no hierarchy.
  - When leaves and containers have drastically different behaviors that are hard to unify.
  - When you need to distinguish leaves from containers at compile time.

Tips and pitfalls:
  - Watch out for cycles -- a directory could accidentally contain itself, causing
    infinite recursion in GetSize() and Display().
  - Composite vs. Decorator: Composite builds hierarchical trees, Decorator adds
    behaviors to a single object.
  - strings.Builder is the idiomatic and efficient way to build strings in Go.
*/
package composite

import (
	"fmt"
	"strings"
)

// Component defines the common interface for files and directories in the composite tree.
type Component interface {
	GetName() string
	GetSize() int
	Display(indent string) string
}

// File represents a file -- a terminal element (leaf) of the tree that cannot contain children.
type File struct {
	name string
	size int
}

// NewFile creates a new file with the given name and size.
func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

// GetName returns the file name.
func (f *File) GetName() string {
	return f.name
}

// GetSize returns the file size in bytes.
func (f *File) GetSize() int {
	return f.size
}

// Display returns a text representation of the file with the given indentation.
func (f *File) Display(indent string) string {
	return fmt.Sprintf("%s%s (%d bytes)", indent, f.name, f.size)
}

// Directory represents a directory -- a container (composite) that can hold files and other directories.
type Directory struct {
	name     string
	children []Component
}

// NewDirectory creates a new directory with the given name.
func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

// GetName returns the directory name.
func (d *Directory) GetName() string {
	return d.name
}

// GetSize recursively sums the sizes of all children in the directory.
func (d *Directory) GetSize() int {
	total := 0
	for _, child := range d.children {
		total += child.GetSize()
	}
	return total
}

// Display recursively renders the directory tree with indentation.
func (d *Directory) Display(indent string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s%s/", indent, d.name))
	for _, child := range d.children {
		b.WriteString("\n")
		b.WriteString(child.Display(indent + "  "))
	}
	return b.String()
}

// Add adds a component (file or directory) as a child of this directory.
func (d *Directory) Add(c Component) {
	d.children = append(d.children, c)
}
