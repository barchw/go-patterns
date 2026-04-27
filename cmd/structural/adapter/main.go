package main

import (
	"fmt"

	"claude-test/internal/structural/adapter"
)

func main() {
	fmt.Println("=== Adapter Pattern ===")

	xmlData := &adapter.XMLData{
		Content: `<person><name>Alice</name><age>30</age><city>Berlin</city></person>`,
	}

	jsonAdapter := adapter.NewXMLToJSONAdapter(xmlData)
	fmt.Println("XML input:", xmlData.Content)
	fmt.Println("JSON output:", jsonAdapter.GetJSON())
}
