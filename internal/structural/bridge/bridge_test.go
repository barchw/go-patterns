package bridge

import "testing"

func TestBridge(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected string
	}{
		{
			name:     "circle with vector renderer",
			shape:    NewCircle(&VectorRenderer{}, 5.0),
			expected: "Drawing circle as vector lines with radius 5.0",
		},
		{
			name:     "circle with raster renderer",
			shape:    NewCircle(&RasterRenderer{}, 3.5),
			expected: "Drawing circle as pixels with radius 3.5",
		},
		{
			name:     "square with vector renderer",
			shape:    NewSquare(&VectorRenderer{}, 4.0),
			expected: "Drawing square as vector lines with side 4.0",
		},
		{
			name:     "square with raster renderer",
			shape:    NewSquare(&RasterRenderer{}, 7.2),
			expected: "Drawing square as pixels with side 7.2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Draw()
			if got != tt.expected {
				t.Fatalf("got %q, want %q", got, tt.expected)
			}
		})
	}
}
