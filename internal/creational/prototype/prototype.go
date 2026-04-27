/*
Package prototype implements the Prototype pattern.

What is it?
Prototype allows creating new objects by cloning existing instances instead of
constructing them from scratch. The key point is performing a deep copy so that
the clone and the original do not share references to the same data.

When to use?
  - When creating an object from scratch is expensive and cloning is significantly cheaper.
  - When you need many variants of an object that differ only in a few fields.
  - When you want to avoid a hierarchy of factories -- you store pre-configured prototypes instead.
  - When the exact type of the object is not known at compile time (you operate on an interface).

When NOT to use?
  - When objects are simple and have no nested reference structures.
  - When a deep copy is hard to implement (cyclic references, file handles).
  - When every new object is unique and no common template can be extracted.

Tips and pitfalls:
  - Deep copy is critical -- a simple *d operation copies only pointers, not the data.
  - Do not clone objects containing sync.Mutex, sync.WaitGroup, etc. -- they require fresh instances.
  - Consider a prototype registry (map[string]Prototype) in larger systems.
  - A type assertion (Clone().(*DocumentTemplate)) is necessary because Clone() returns an interface.
*/
package prototype

// Prototype defines the cloning interface -- every cloneable type must implement the Clone() method.
type Prototype interface {
	Clone() Prototype
}

// DocumentTemplate represents a document template with value fields and a reference field (Metadata).
type DocumentTemplate struct {
	Title    string
	Body     string
	Metadata map[string]string
}

// NewDocumentTemplate creates a new document template, making a defensive copy of the provided metadata map.
func NewDocumentTemplate(title, body string, metadata map[string]string) *DocumentTemplate {
	md := make(map[string]string, len(metadata))
	for k, v := range metadata {
		md[k] = v
	}
	return &DocumentTemplate{
		Title:    title,
		Body:     body,
		Metadata: md,
	}
}

// Clone creates a deep copy of the document template, ensuring the clone is independent from the original.
func (d *DocumentTemplate) Clone() Prototype {
	md := make(map[string]string, len(d.Metadata))
	for k, v := range d.Metadata {
		md[k] = v
	}
	return &DocumentTemplate{
		Title:    d.Title,
		Body:     d.Body,
		Metadata: md,
	}
}
