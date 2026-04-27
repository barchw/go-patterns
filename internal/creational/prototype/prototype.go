package prototype

type Prototype interface {
	Clone() Prototype
}

type DocumentTemplate struct {
	Title    string
	Body     string
	Metadata map[string]string
}

func NewDocumentTemplate(title, body string, metadata map[string]string) *DocumentTemplate {
	md := make(map[string]string, len(metadata))
	for k, v := range metadata {
		md[k] = v
	}
	return &DocumentTemplate{
		Title:    title,
		Body:     body,
		Metadata: md,
	}
}

func (d *DocumentTemplate) Clone() Prototype {
	md := make(map[string]string, len(d.Metadata))
	for k, v := range d.Metadata {
		md[k] = v
	}
	return &DocumentTemplate{
		Title:    d.Title,
		Body:     d.Body,
		Metadata: md,
	}
}
