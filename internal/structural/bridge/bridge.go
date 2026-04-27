/*
Package bridge implements the Bridge design pattern.

What is it?
Bridge separates an abstraction from its implementation, so that both can change
independently. Instead of building class hierarchies through inheritance, the pattern
connects the abstraction with the implementation through composition. In this package,
the abstraction is shapes (Shape), and the implementation is rendering methods (Renderer).

When to use?
  - When you want to avoid a combinatorial explosion of subclasses (e.g., 3 shapes x 2 renderers).
  - When the abstraction and implementation should be able to change independently at runtime.
  - When you want to be able to swap the implementation while the program is running.
  - When you're working with multiple platforms or backends and want to share the abstraction logic.

When NOT to use?
  - When you have only one implementation and don't plan to add more.
  - When the hierarchy is flat and simple -- an extra layer of abstraction makes the code harder to understand.
  - When the abstraction and implementation are tightly coupled and there's no point in separating them.

Tips and pitfalls:
  - Bridge vs. Strategy: both use composition, but Bridge separates two independent hierarchies,
    while Strategy allows swapping a single algorithm.
  - Avoid "fat" interfaces -- when adding a new shape you need to modify
    the Renderer interface and all of its implementations.
  - Renderer is injected through the constructor -- that is the "bridge" created from the outside.
*/
package bridge

import "fmt"

// Renderer defines the implementation interface for rendering shapes.
type Renderer interface {
	RenderCircle(radius float64) string
	RenderSquare(side float64) string
}

// VectorRenderer renders shapes as vector lines.
type VectorRenderer struct{}

// RenderCircle draws a circle as vector lines with the given radius.
func (v *VectorRenderer) RenderCircle(radius float64) string {
	return fmt.Sprintf("Drawing circle as vector lines with radius %.1f", radius)
}

// RenderSquare draws a square as vector lines with the given side length.
func (v *VectorRenderer) RenderSquare(side float64) string {
	return fmt.Sprintf("Drawing square as vector lines with side %.1f", side)
}

// RasterRenderer renders shapes as pixels (raster).
type RasterRenderer struct{}

// RenderCircle draws a circle as pixels with the given radius.
func (r *RasterRenderer) RenderCircle(radius float64) string {
	return fmt.Sprintf("Drawing circle as pixels with radius %.1f", radius)
}

// RenderSquare draws a square as pixels with the given side length.
func (r *RasterRenderer) RenderSquare(side float64) string {
	return fmt.Sprintf("Drawing square as pixels with side %.1f", side)
}

// Shape defines the abstraction interface for a shape with a Draw method.
type Shape interface {
	Draw() string
}

// Circle represents a circle that delegates drawing to a Renderer.
type Circle struct {
	renderer Renderer
	radius   float64
}

// NewCircle creates a new circle with the given renderer and radius.
func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

// Draw draws the circle using the assigned renderer.
func (c *Circle) Draw() string {
	return c.renderer.RenderCircle(c.radius)
}

// Square represents a square that delegates drawing to a Renderer.
type Square struct {
	renderer Renderer
	side     float64
}

// NewSquare creates a new square with the given renderer and side length.
func NewSquare(renderer Renderer, side float64) *Square {
	return &Square{renderer: renderer, side: side}
}

// Draw draws the square using the assigned renderer.
func (s *Square) Draw() string {
	return s.renderer.RenderSquare(s.side)
}
