package main

import (
	"fmt"

	"claude-test/internal/behavioral/strategy"
)

func main() {
	fmt.Println("=== Strategy Pattern ===")

	data := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Input: %v\n", data)

	sorter := strategy.NewSorter(strategy.BubbleSort)
	fmt.Printf("BubbleSort:    %v\n", sorter.Sort(data))

	sorter.SetStrategy(strategy.InsertionSort)
	fmt.Printf("InsertionSort: %v\n", sorter.Sort(data))

	sorter.SetStrategy(strategy.QuickSort)
	fmt.Printf("QuickSort:     %v\n", sorter.Sort(data))
}
