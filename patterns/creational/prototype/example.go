package prototype

type Document struct {
	Title   string
	Content string
	Tags    []string
}

func (d *Document) Clone() *Document {
	// Deep copy Tags
	tagsCopy := make([]string, len(d.Tags))
	copy(tagsCopy, d.Tags)
	return &Document{
		Title:   d.Title,
		Content: d.Content,
		Tags:    tagsCopy,
	}
}