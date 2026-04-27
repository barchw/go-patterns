/*
Package strategy -- Strategy

What is it?
Strategy is a behavioral design pattern that lets you define a family of algorithms,
encapsulate each one in a separate struct (or function), and make them interchangeable. The client can
dynamically swap the algorithm at runtime without modifying the code that uses it.

When to use?
  - When you have several variants of the same algorithm and want to easily swap them at runtime.
  - When you want to avoid sprawling if/else or switch statements that select an algorithm.
  - When different contexts require different behaviors, but the interface remains the same.
  - When you want to add new algorithms without modifying existing code (Open/Closed principle).

When NOT to use?
  - When you have only one algorithm and don't anticipate more -- the extra abstraction complicates the code.
  - When the differences between algorithms are minimal -- a simple control parameter is better.
  - When clients don't need to know about different strategies -- if the algorithm choice is fixed, the pattern is unnecessary.

Tips and pitfalls:
  - Function types in Go: if a strategy has only one method, use a function type instead of an interface.
  - Data copying: each strategy creates a copy of the input data, which prevents side effects.
  - Nil guard: always check whether a strategy has been set.
  - Strategy vs State: Strategy lets the client choose the algorithm; State changes behavior automatically.
*/
package strategy

// SortStrategy defines a function type representing a sorting strategy.
type SortStrategy func([]int) []int

// Sorter is the context that holds the current sorting strategy.
type Sorter struct {
	strategy SortStrategy
}

// NewSorter creates a new Sorter with the given sorting strategy.
func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

// SetStrategy changes the current sorting strategy at runtime.
func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

// Sort sorts the data using the current strategy. Returns the data unchanged if the strategy is nil.
func (s *Sorter) Sort(data []int) []int {
	if s.strategy == nil {
		return data
	}
	return s.strategy(data)
}

// BubbleSort implements the bubble sort strategy. Creates a copy of the input data.
func BubbleSort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	n := len(result)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// InsertionSort implements the insertion sort strategy. Creates a copy of the input data.
func InsertionSort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}

// QuickSort implements the quick sort strategy. Creates a copy of the input data.
func QuickSort(data []int) []int {
	result := make([]int, len(data))
	copy(result, data)
	quickSortHelper(result, 0, len(result)-1)
	return result
}

func quickSortHelper(data []int, low, high int) {
	if low < high {
		pi := partition(data, low, high)
		quickSortHelper(data, low, pi-1)
		quickSortHelper(data, pi+1, high)
	}
}

func partition(data []int, low, high int) int {
	pivot := data[high]
	i := low - 1
	for j := low; j < high; j++ {
		if data[j] <= pivot {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[i+1], data[high] = data[high], data[i+1]
	return i + 1
}
