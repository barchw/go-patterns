package main

import (
	"fmt"

	"claude-test/internal/structural/bridge"
)

func main() {
	fmt.Println("=== Bridge Pattern ===")

	vector := &bridge.VectorRenderer{}
	raster := &bridge.RasterRenderer{}

	shapes := []bridge.Shape{
		bridge.NewCircle(vector, 5.0),
		bridge.NewCircle(raster, 5.0),
		bridge.NewSquare(vector, 3.0),
		bridge.NewSquare(raster, 3.0),
	}

	for _, s := range shapes {
		fmt.Println(s.Draw())
	}
}
