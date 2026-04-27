package main

import (
	"fmt"

	"claude-test/internal/behavioral/template"
)

func main() {
	fmt.Println("=== Template Method Pattern ===")

	data := []map[string]string{
		{"name": "Alice", "age": "30", "city": "NYC"},
		{"name": "Bob", "age": "25", "city": "LA"},
		{"name": "Charlie", "age": "35", "city": "Chicago"},
	}

	fmt.Println("CSV Export:")
	csv := template.NewCSVExporter()
	fmt.Println(csv.Export(data))

	fmt.Println("JSON Export:")
	json := template.NewJSONExporter()
	fmt.Println(json.Export(data))

	fmt.Println("XML Export:")
	xml := template.NewXMLExporter("users", "user")
	fmt.Println(xml.Export(data))
}
