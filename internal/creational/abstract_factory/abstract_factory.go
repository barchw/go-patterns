package abstract_factory

import "fmt"

type Button interface {
	Render() string
}

type Checkbox interface {
	Render() string
}

type UIFactory interface {
	CreateButton(label string) Button
	CreateCheckbox(label string) Checkbox
}

type lightButton struct {
	label string
}

func (b *lightButton) Render() string {
	return fmt.Sprintf("[Light Button: %s]", b.label)
}

type lightCheckbox struct {
	label string
}

func (c *lightCheckbox) Render() string {
	return fmt.Sprintf("[Light Checkbox: %s]", c.label)
}

type LightThemeFactory struct{}

func (f *LightThemeFactory) CreateButton(label string) Button {
	return &lightButton{label: label}
}

func (f *LightThemeFactory) CreateCheckbox(label string) Checkbox {
	return &lightCheckbox{label: label}
}

type darkButton struct {
	label string
}

func (b *darkButton) Render() string {
	return fmt.Sprintf("[Dark Button: %s]", b.label)
}

type darkCheckbox struct {
	label string
}

func (c *darkCheckbox) Render() string {
	return fmt.Sprintf("[Dark Checkbox: %s]", c.label)
}

type DarkThemeFactory struct{}

func (f *DarkThemeFactory) CreateButton(label string) Button {
	return &darkButton{label: label}
}

func (f *DarkThemeFactory) CreateCheckbox(label string) Checkbox {
	return &darkCheckbox{label: label}
}
