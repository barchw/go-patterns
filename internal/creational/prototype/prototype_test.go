package prototype

import "testing"

func TestDocumentTemplate_Clone_DeepCopy(t *testing.T) {
	original := NewDocumentTemplate(
		"Report",
		"Quarterly results",
		map[string]string{
			"author":  "Alice",
			"version": "1.0",
		},
	)

	cloned := original.Clone().(*DocumentTemplate)

	if cloned.Title != original.Title {
		t.Errorf("cloned Title = %q, want %q", cloned.Title, original.Title)
	}
	if cloned.Body != original.Body {
		t.Errorf("cloned Body = %q, want %q", cloned.Body, original.Body)
	}
	if cloned.Metadata["author"] != "Alice" {
		t.Errorf("cloned Metadata[author] = %q, want %q", cloned.Metadata["author"], "Alice")
	}

	cloned.Title = "Modified Report"
	cloned.Body = "Updated content"
	cloned.Metadata["author"] = "Bob"
	cloned.Metadata["reviewer"] = "Charlie"

	if original.Title != "Report" {
		t.Errorf("original Title changed to %q after modifying clone", original.Title)
	}
	if original.Body != "Quarterly results" {
		t.Errorf("original Body changed to %q after modifying clone", original.Body)
	}
	if original.Metadata["author"] != "Alice" {
		t.Errorf("original Metadata[author] changed to %q after modifying clone", original.Metadata["author"])
	}
	if _, exists := original.Metadata["reviewer"]; exists {
		t.Error("original Metadata gained 'reviewer' key after modifying clone")
	}
}
