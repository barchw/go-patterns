package decorator

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type DataSource interface {
	WriteData(data string) string
	ReadData() string
}

type FileDataSource struct {
	filename string
	data     string
}

func NewFileDataSource(filename string) *FileDataSource {
	return &FileDataSource{filename: filename}
}

func (f *FileDataSource) WriteData(data string) string {
	f.data = data
	return fmt.Sprintf("[FileDataSource:%s] wrote %d bytes", f.filename, len(data))
}

func (f *FileDataSource) ReadData() string {
	return f.data
}

type CompressionDecorator struct {
	wrapped DataSource
}

func NewCompressionDecorator(wrapped DataSource) *CompressionDecorator {
	return &CompressionDecorator{wrapped: wrapped}
}

func (c *CompressionDecorator) WriteData(data string) string {
	compressed := compress(data)
	result := c.wrapped.WriteData(compressed)
	return fmt.Sprintf("[Compressed] %s", result)
}

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

type EncryptionDecorator struct {
	wrapped DataSource
}

func NewEncryptionDecorator(wrapped DataSource) *EncryptionDecorator {
	return &EncryptionDecorator{wrapped: wrapped}
}

func (e *EncryptionDecorator) WriteData(data string) string {
	encrypted := encrypt(data)
	result := e.wrapped.WriteData(encrypted)
	return fmt.Sprintf("[Encrypted] %s", result)
}

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
