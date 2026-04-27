package composite

import (
	"fmt"
	"strings"
)

type Component interface {
	GetName() string
	GetSize() int
	Display(indent string) string
}

type File struct {
	name string
	size int
}

func NewFile(name string, size int) *File {
	return &File{name: name, size: size}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetSize() int {
	return f.size
}

func (f *File) Display(indent string) string {
	return fmt.Sprintf("%s%s (%d bytes)", indent, f.name, f.size)
}

type Directory struct {
	name     string
	children []Component
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) GetName() string {
	return d.name
}

func (d *Directory) GetSize() int {
	total := 0
	for _, child := range d.children {
		total += child.GetSize()
	}
	return total
}

func (d *Directory) Display(indent string) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("%s%s/", indent, d.name))
	for _, child := range d.children {
		b.WriteString("\n")
		b.WriteString(child.Display(indent + "  "))
	}
	return b.String()
}

func (d *Directory) Add(c Component) {
	d.children = append(d.children, c)
}
