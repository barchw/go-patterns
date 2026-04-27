package main

import (
	"fmt"

	"claude-test/internal/structural/proxy"
)

func main() {
	fmt.Println("=== Proxy Pattern ===")

	real := proxy.NewRealWeatherService()
	cached := proxy.NewCachedWeatherService(real)

	fmt.Println("First call (Berlin):", cached.GetWeather("Berlin"))
	fmt.Println("Second call (Berlin):", cached.GetWeather("Berlin"))
	fmt.Println("First call (Tokyo):", cached.GetWeather("Tokyo"))

	fmt.Printf("\nReal service was called %d times\n", real.CallCount())
}
