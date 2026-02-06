package main

import "fmt"
import "errors"

func SliceCopy(nums []int) []int {
	numsCopy := make([]int, len(nums))
	copy(numsCopy, nums)
	return numsCopy
}

func Join1(nums1, nums2 []int) []int {
	nums1Copy := SliceCopy(nums1)
	nums1Copy = append(nums1Copy, nums2...)
	return nums1Copy
}

func Join(nums1, nums2 []int) []int {
    totalLen := len(nums1) + len(nums2)
    result := make([]int, totalLen, totalLen)
    copy(result, nums1)
    copy(result[len(nums1):], nums2)
    return result
}


func Mix(nums []int) []int	{
	diff := len(nums) / 2
	numsCopy := make([]int, len(nums))	
	i := 0
	j := 0
	for i < len(nums) / 2 {
		numsCopy[j] = nums[i]
		numsCopy[j + 1] = nums[i + diff]
		i += 1
		j += 2
	}
	return numsCopy
}




func UnderLimit(nums []int, limit int, n int) ([]int, error) {
	if nums == nil {
		return nil, errors.New("nums не может быть nil")
	}
	if n < 0 {
		return nil, errors.New("n не может быть отрицательным")
	}
	var result []int
	for _, price := range nums {
		if price < limit {
			result = append(result, price)
			if len(result) == n {
				break
			}
		}
	}
	return result, nil
}



func main() {
	nums := []int{}
	n := -1
	limit := -5

	// nums := []int{-13, 0, 6}
	// n := 1
	// limit := -5

	// nums := []int{3, 5, 6}
	// n := 5
	// limit := 10

	// nums := []int{}
	// n := 5
	// limit := 3

	// nums := []int{4, 7, 89, 3, 21, 2, 5, 7, 32, 4, 6, 8, 0, 3, 4, 6, 2, 115, 12}
	// n := 5
	// limit := 3

    fmt.Println(UnderLimit(nums, limit, n))
}

