package prototype

import "testing"

func TestDocumentClone(t *testing.T) {
	doc1 := &Document{
		Title:   "Original",
		Content: "This is the original document.",
		Tags:    []string{"go", "patterns"},
	}
	doc2 := doc1.Clone()

	if doc1 == doc2 {
		t.Errorf("Clone should return a different instance")
	}
	if doc1.Title != doc2.Title || doc1.Content != doc2.Content {
		t.Errorf("Cloned document fields do not match")
	}
	if &doc1.Tags[0] == &doc2.Tags[0] {
		t.Errorf("Tags slice should be deep copied")
	}
}