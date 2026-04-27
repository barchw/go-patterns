package adapter

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

type JSONData interface {
	GetJSON() string
}

type XMLData struct {
	Content string
}

type xmlEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
	Entries []xmlEntry `xml:",any"`
}

type XMLToJSONAdapter struct {
	xmlData *XMLData
}

func NewXMLToJSONAdapter(xmlData *XMLData) *XMLToJSONAdapter {
	return &XMLToJSONAdapter{xmlData: xmlData}
}

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
