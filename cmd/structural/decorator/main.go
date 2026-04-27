package main

import (
	"fmt"

	"claude-test/internal/structural/decorator"
)

func main() {
	fmt.Println("=== Decorator Pattern ===")

	base := decorator.NewFileDataSource("data.txt")
	compressed := decorator.NewCompressionDecorator(base)
	encrypted := decorator.NewEncryptionDecorator(compressed)

	input := "Hello, Decorator Pattern!"
	fmt.Println("Input:", input)
	fmt.Println("Write:", encrypted.WriteData(input))
	fmt.Println("Read:", encrypted.ReadData())
}
