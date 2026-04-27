/*
Package decorator implements the Decorator design pattern.

What is it?
Decorator allows dynamically adding new behaviors to objects by wrapping them
in other objects that implement the same interface. It works like "onion layers" -- each
decorator adds its own functionality and delegates the rest to the wrapped object. In this
package, the pattern models a data source (DataSource) with compression and encryption decorators.

When to use?
  - When you want to add behaviors to objects at runtime, without modifying their code.
  - When there are too many behavior combinations for embedding/inheritance.
  - When you want to combine different behaviors in any order.
  - When you follow the Open/Closed principle -- open for extension, closed for modification.

When NOT to use?
  - When the order of decorators matters and errors are hard to detect.
  - When you need access to the internal state of the original object.
  - When a simple if/switch is enough to differentiate behaviors.

Tips and pitfalls:
  - Order matters: Encryption(Compression(File)) is not the same as
    Compression(Encryption(File)).
  - Decorator does not change the interface (it adds behavior), Adapter changes the interface.
  - In Go each decorator must separately implement the entire interface and delegate to wrapped.
  - Each decorator can be tested independently by creating a simple mock DataSource.
*/
package decorator

import (
	"encoding/base64"
	"fmt"
	"strings"
)

// DataSource defines the common interface for data sources and decorators.
type DataSource interface {
	WriteData(data string) string
	ReadData() string
}

// FileDataSource is the base data source that stores data in memory (simulating a file).
type FileDataSource struct {
	filename string
	data     string
}

// NewFileDataSource creates a new data source associated with the given filename.
func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

// WriteData writes data to the source and returns a write confirmation message.
func (f *FileDataSource) WriteData(data string) string {
	f.data = data
	return fmt.Sprintf("[FileDataSource:%s] wrote %d bytes", f.filename, len(data))
}

// ReadData reads data from the source.
func (f *FileDataSource) ReadData() string {
	return f.data
}

// CompressionDecorator wraps a DataSource, adding compression on write and decompression on read.
type CompressionDecorator struct {
	wrapped DataSource
}

// NewCompressionDecorator creates a new compression decorator wrapping the given data source.
func NewCompressionDecorator(wrapped DataSource) *CompressionDecorator {
	return &CompressionDecorator{wrapped: wrapped}
}

// WriteData compresses the data before passing it to the wrapped source.
func (c *CompressionDecorator) WriteData(data string) string {
	compressed := compress(data)
	result := c.wrapped.WriteData(compressed)
	return fmt.Sprintf("[Compressed] %s", result)
}

// ReadData reads data from the wrapped source and decompresses it.
func (c *CompressionDecorator) ReadData() string {
	data := c.wrapped.ReadData()
	return decompress(data)
}

func compress(data string) string {
	return "COMPRESSED:" + strings.ReplaceAll(data, " ", "_")
}

func decompress(data string) string {
	if strings.HasPrefix(data, "COMPRESSED:") {
		return strings.ReplaceAll(strings.TrimPrefix(data, "COMPRESSED:"), "_", " ")
	}
	return data
}

// EncryptionDecorator wraps a DataSource, adding encryption (base64) on write and decryption on read.
type EncryptionDecorator struct {
	wrapped DataSource
}

// NewEncryptionDecorator creates a new encryption decorator wrapping the given data source.
func NewEncryptionDecorator(wrapped DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{wrapped: wrapped}
}

// WriteData encrypts the data before passing it to the wrapped source.
func (e *EncryptionDecorator) WriteData(data string) string {
	encrypted := encrypt(data)
	result := e.wrapped.WriteData(encrypted)
	return fmt.Sprintf("[Encrypted] %s", result)
}

// ReadData reads data from the wrapped source and decrypts it.
func (e *EncryptionDecorator) ReadData() string {
	data := e.wrapped.ReadData()
	return decrypt(data)
}

func encrypt(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func decrypt(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return data
	}
	return string(decoded)
}
