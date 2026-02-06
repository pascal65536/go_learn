package main

import (
	"fmt"
	"sort"
)

func SortAndMerge(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	left := []int{5, 2, 8, 1}
	right := []int{7, 3, 6, 4}
	merged := SortAndMerge(left, right)
	fmt.Println(merged) // [1 2 3 4 5 6 7 8]
}
