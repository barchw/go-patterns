/*
Package abstract_factory implements the Abstract Factory pattern.

What is it?
Abstract Factory provides an interface for creating families of related objects
without specifying their concrete classes. Unlike the Factory Method, which creates
a single product, Abstract Factory coordinates the creation of multiple products,
guaranteeing their consistency (e.g., all UI elements in the same theme).

When to use?
  - When the system must create families of related objects (e.g., UI elements in a specific theme).
  - When you want to guarantee that products from the same family are used together.
  - When the product set may grow with new families without modifying client code.
  - When client code should operate solely on interfaces.

When NOT to use?
  - When you have only one product type -- a simple Factory Method is enough.
  - When adding a new product type is more likely than adding a new family.
  - When the complexity of the interface hierarchy is not justified by the project's scale.

Tips and pitfalls:
  - Adding a new family (e.g., HighContrastThemeFactory) is cheap -- just a new struct.
  - Adding a new product type (e.g., Slider) is expensive -- it requires changes in all factories.
  - Export factories, hide products -- the client should never create a lightButton directly.
  - Pass factories as function parameters, not as global variables -- this makes testing easier.
*/
package abstract_factory

import "fmt"

// Button defines the UI button interface with a render method.
type Button interface {
	Render() string
}

// Checkbox defines the UI checkbox interface with a render method.
type Checkbox interface {
	Render() string
}

// UIFactory defines the abstract factory interface that groups the creation of all UI elements into a single family.
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

// LightThemeFactory is a concrete factory that creates UI elements in the light theme.
type LightThemeFactory struct{}

// CreateButton creates a button in the light theme.
func (f *LightThemeFactory) CreateButton(label string) Button {
	return &lightButton{label: label}
}

// CreateCheckbox creates a checkbox in the light theme.
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

// DarkThemeFactory is a concrete factory that creates UI elements in the dark theme.
type DarkThemeFactory struct{}

// CreateButton creates a button in the dark theme.
func (f *DarkThemeFactory) CreateButton(label string) Button {
	return &darkButton{label: label}
}

// CreateCheckbox creates a checkbox in the dark theme.
func (f *DarkThemeFactory) CreateCheckbox(label string) Checkbox {
	return &darkCheckbox{label: label}
}
