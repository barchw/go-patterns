/*
Package adapter implements the Adapter design pattern.

What is it?
Adapter is a structural design pattern that allows objects with incompatible interfaces
to work together. It acts as a "translator" -- it wraps an existing object and exposes
it through a new, expected interface. In this package, the Adapter converts XML data
to JSON format.

When to use?
  - When you want to use an existing structure, but its interface doesn't fit the rest of your code.
  - When you're integrating an external library with a different API than your system expects.
  - When you need an intermediary layer between two independent modules.
  - When you're migrating from one data format to another (e.g., XML to JSON).

When NOT to use?
  - When you can simply change the source object's interface -- the Adapter adds unnecessary complexity.
  - When you're adapting many different interfaces at the same time -- this leads to hard-to-maintain code.
  - When the difference between interfaces requires complex business logic in the adapter.

Tips and pitfalls:
  - Adapter changes the interface of a single object; Facade simplifies the interface of an entire subsystem -- don't confuse these patterns.
  - In Go you don't need "implements" -- interface satisfaction is implicit (duck typing).
  - Avoid "super-adapters" that translate many interfaces; it's better to have several small, dedicated ones.
*/
package adapter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

// JSONData defines the target interface that returns data in JSON format.
type JSONData interface {
	GetJSON() string
}

// XMLData stores raw XML data as a string (the adaptee type).
type XMLData struct {
	Content string
}

type xmlEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
	Entries []xmlEntry `xml:",any"`
}

// XMLToJSONAdapter wraps XMLData and implements the JSONData interface,
// converting XML data to JSON format.
type XMLToJSONAdapter struct {
	xmlData *XMLData
}

// NewXMLToJSONAdapter creates a new adapter that converts XMLData to JSONData.
func NewXMLToJSONAdapter(xmlData *XMLData) *XMLToJSONAdapter {
	return &XMLToJSONAdapter{xmlData: xmlData}
}

// GetJSON parses the XML data from the wrapped XMLData and returns it as a JSON string.
func (a *XMLToJSONAdapter) GetJSON() string {
	content := strings.TrimSpace(a.xmlData.Content)
	if content == "" {
		return "{}"
	}

	var entry xmlEntry
	if err := xml.Unmarshal([]byte(content), &entry); err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err.Error())
	}

	result := entryToMap(entry)
	data, err := json.Marshal(result)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}`, err.Error())
	}
	return string(data)
}

func entryToMap(e xmlEntry) map[string]interface{} {
	m := make(map[string]interface{})
	if len(e.Entries) == 0 {
		m[e.XMLName.Local] = e.Value
		return m
	}

	children := make(map[string]interface{})
	for _, child := range e.Entries {
		if len(child.Entries) == 0 {
			children[child.XMLName.Local] = child.Value
		} else {
			children[child.XMLName.Local] = flattenEntry(child)
		}
	}
	m[e.XMLName.Local] = children
	return m
}

func flattenEntry(e xmlEntry) map[string]interface{} {
	children := make(map[string]interface{})
	for _, child := range e.Entries {
		if len(child.Entries) == 0 {
			children[child.XMLName.Local] = child.Value
		} else {
			children[child.XMLName.Local] = flattenEntry(child)
		}
	}
	return children
}
