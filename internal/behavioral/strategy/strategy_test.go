package strategy

import (
	"reflect"
	"testing"
)

func TestSorter(t *testing.T) {
	tests := []struct {
		name     string
		strategy SortStrategy
		input    []int
		expected []int
	}{
		{
			name:     "BubbleSort sorts unsorted slice",
			strategy: BubbleSort,
			input:    []int{5, 3, 8, 1, 2},
			expected: []int{1, 2, 3, 5, 8},
		},
		{
			name:     "InsertionSort sorts unsorted slice",
			strategy: InsertionSort,
			input:    []int{5, 3, 8, 1, 2},
			expected: []int{1, 2, 3, 5, 8},
		},
		{
			name:     "QuickSort sorts unsorted slice",
			strategy: QuickSort,
			input:    []int{5, 3, 8, 1, 2},
			expected: []int{1, 2, 3, 5, 8},
		},
		{
			name:     "BubbleSort handles already sorted",
			strategy: BubbleSort,
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "InsertionSort handles reverse sorted",
			strategy: InsertionSort,
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "QuickSort handles single element",
			strategy: QuickSort,
			input:    []int{42},
			expected: []int{42},
		},
		{
			name:     "BubbleSort handles empty slice",
			strategy: BubbleSort,
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "QuickSort handles duplicates",
			strategy: QuickSort,
			input:    []int{3, 1, 3, 1, 2},
			expected: []int{1, 1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorter := NewSorter(tt.strategy)
			result := sorter.Sort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSorterSetStrategy(t *testing.T) {
	sorter := NewSorter(BubbleSort)
	input := []int{5, 3, 1}
	expected := []int{1, 3, 5}

	result := sorter.Sort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort: got %v, want %v", result, expected)
	}

	sorter.SetStrategy(QuickSort)
	result = sorter.Sort(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("QuickSort after switch: got %v, want %v", result, expected)
	}
}

func TestSorterDoesNotMutateInput(t *testing.T) {
	input := []int{5, 3, 8, 1, 2}
	original := make([]int, len(input))
	copy(original, input)

	sorter := NewSorter(QuickSort)
	sorter.Sort(input)

	if !reflect.DeepEqual(input, original) {
		t.Errorf("input was mutated: got %v, want %v", input, original)
	}
}
