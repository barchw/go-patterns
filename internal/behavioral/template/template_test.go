package template

import (
	"strings"
	"testing"
)

func TestExporters(t *testing.T) {
	data := []map[string]string{
		{"name": "Alice", "age": "30"},
		{"name": "Bob", "age": "25"},
	}

	tests := []struct {
		name     string
		exporter Exporter
		contains []string
	}{
		{
			name:     "CSVExporter",
			exporter: NewCSVExporter(),
			contains: []string{
				"age,name\n",
				"30,Alice\n",
				"25,Bob\n",
			},
		},
		{
			name:     "JSONExporter",
			exporter: NewJSONExporter(),
			contains: []string{
				"[\n",
				"\"name\": \"Alice\"",
				"\"age\": \"30\"",
				"\"name\": \"Bob\"",
				"\"age\": \"25\"",
				"]\n",
			},
		},
		{
			name:     "XMLExporter",
			exporter: NewXMLExporter("people", "person"),
			contains: []string{
				"<people>\n",
				"<person>",
				"<name>Alice</name>",
				"<age>30</age>",
				"<name>Bob</name>",
				"<age>25</age>",
				"</people>\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.exporter.Export(data)
			for _, substr := range tt.contains {
				if !strings.Contains(result, substr) {
					t.Errorf("output missing %q\ngot:\n%s", substr, result)
				}
			}
		})
	}
}

func TestCSVExporter_Format(t *testing.T) {
	data := []map[string]string{
		{"x": "1", "y": "2"},
	}
	exporter := NewCSVExporter()
	result := exporter.Export(data)
	expected := "x,y\n1,2\n"
	if result != expected {
		t.Errorf("got:\n%s\nwant:\n%s", result, expected)
	}
}

func TestJSONExporter_Format(t *testing.T) {
	data := []map[string]string{
		{"key": "val"},
	}
	exporter := NewJSONExporter()
	result := exporter.Export(data)
	if !strings.HasPrefix(result, "[\n") {
		t.Errorf("should start with [\\n, got: %q", result[:5])
	}
	if !strings.HasSuffix(result, "]\n") {
		t.Errorf("should end with ]\\n, got: %q", result[len(result)-5:])
	}
}

func TestXMLExporter_Format(t *testing.T) {
	data := []map[string]string{
		{"id": "1"},
	}
	exporter := NewXMLExporter("items", "item")
	result := exporter.Export(data)
	if !strings.HasPrefix(result, "<items>\n") {
		t.Errorf("should start with <items>, got: %q", result)
	}
	if !strings.HasSuffix(result, "</items>\n") {
		t.Errorf("should end with </items>, got: %q", result)
	}
	if !strings.Contains(result, "<item>") || !strings.Contains(result, "</item>") {
		t.Errorf("should contain <item> tags, got: %q", result)
	}
}

func TestExporter_EmptyData(t *testing.T) {
	exporters := []struct {
		name     string
		exporter Exporter
	}{
		{"CSV", NewCSVExporter()},
		{"JSON", NewJSONExporter()},
		{"XML", NewXMLExporter("root", "row")},
	}

	for _, tt := range exporters {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.exporter.Export(nil)
			if result != "" {
				t.Errorf("expected empty string for nil data, got %q", result)
			}

			result = tt.exporter.Export([]map[string]string{})
			if result != "" {
				t.Errorf("expected empty string for empty data, got %q", result)
			}
		})
	}
}
