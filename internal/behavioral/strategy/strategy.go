package strategy

type SortStrategy func([]int) []int

type Sorter struct {
	strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
	return &Sorter{strategy: strategy}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
	s.strategy = strategy
}

func (s *Sorter) Sort(data []int) []int {
	if s.strategy == nil {
		return data
	}
	return s.strategy(data)
}

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
