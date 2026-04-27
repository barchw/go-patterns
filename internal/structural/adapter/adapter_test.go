package adapter

import (
	"encoding/json"
	"testing"
)

func TestXMLToJSONAdapter(t *testing.T) {
	tests := []struct {
		name     string
		xml      string
		wantKeys map[string]interface{}
		wantErr  bool
	}{
		{
			name: "simple element",
			xml:  `<name>Alice</name>`,
			wantKeys: map[string]interface{}{
				"name": "Alice",
			},
		},
		{
			name: "nested elements",
			xml:  `<person><name>Bob</name><age>30</age></person>`,
			wantKeys: map[string]interface{}{
				"person": map[string]interface{}{
					"name": "Bob",
					"age":  "30",
				},
			},
		},
		{
			name: "deeply nested",
			xml:  `<company><employee><name>Carol</name><role>Engineer</role></employee></company>`,
			wantKeys: map[string]interface{}{
				"company": map[string]interface{}{
					"employee": map[string]interface{}{
						"name": "Carol",
						"role": "Engineer",
					},
				},
			},
		},
		{
			name:     "empty input",
			xml:      "",
			wantKeys: nil,
		},
		{
			name:    "invalid xml",
			xml:     "<broken>",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xmlData := &XMLData{Content: tt.xml}
			var adapter JSONData = NewXMLToJSONAdapter(xmlData)
			result := adapter.GetJSON()

			if tt.xml == "" {
				if result != "{}" {
					t.Fatalf("expected {} for empty input, got %s", result)
				}
				return
			}

			var parsed map[string]interface{}
			if err := json.Unmarshal([]byte(result), &parsed); err != nil {
				t.Fatalf("result is not valid JSON: %s", result)
			}

			if tt.wantErr {
				if _, ok := parsed["error"]; !ok {
					t.Fatal("expected error key in output")
				}
				return
			}

			for key, want := range tt.wantKeys {
				got, ok := parsed[key]
				if !ok {
					t.Fatalf("missing key %q in output: %s", key, result)
				}
				assertDeepEqual(t, key, want, got)
			}
		})
	}
}

func assertDeepEqual(t *testing.T, path string, want, got interface{}) {
	t.Helper()
	switch w := want.(type) {
	case string:
		g, ok := got.(string)
		if !ok {
			t.Fatalf("at %s: expected string, got %T", path, got)
		}
		if w != g {
			t.Fatalf("at %s: want %q, got %q", path, w, g)
		}
	case map[string]interface{}:
		g, ok := got.(map[string]interface{})
		if !ok {
			t.Fatalf("at %s: expected map, got %T", path, got)
		}
		for k, v := range w {
			gv, ok := g[k]
			if !ok {
				t.Fatalf("at %s: missing key %q", path, k)
			}
			assertDeepEqual(t, path+"."+k, v, gv)
		}
	}
}
