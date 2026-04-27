package bridge

import "fmt"

type Renderer interface {
	RenderCircle(radius float64) string
	RenderSquare(side float64) string
}

type VectorRenderer struct{}

func (v *VectorRenderer) RenderCircle(radius float64) string {
	return fmt.Sprintf("Drawing circle as vector lines with radius %.1f", radius)
}

func (v *VectorRenderer) RenderSquare(side float64) string {
	return fmt.Sprintf("Drawing square as vector lines with side %.1f", side)
}

type RasterRenderer struct{}

func (r *RasterRenderer) RenderCircle(radius float64) string {
	return fmt.Sprintf("Drawing circle as pixels with radius %.1f", radius)
}

func (r *RasterRenderer) RenderSquare(side float64) string {
	return fmt.Sprintf("Drawing square as pixels with side %.1f", side)
}

type Shape interface {
	Draw() string
}

type Circle struct {
	renderer Renderer
	radius   float64
}

func NewCircle(renderer Renderer, radius float64) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func (c *Circle) Draw() string {
	return c.renderer.RenderCircle(c.radius)
}

type Square struct {
	renderer Renderer
	side     float64
}

func NewSquare(renderer Renderer, side float64) *Square {
	return &Square{renderer: renderer, side: side}
}

func (s *Square) Draw() string {
	return s.renderer.RenderSquare(s.side)
}
