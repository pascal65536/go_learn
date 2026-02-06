package main

import "fmt"


func FindMaxKey(m map[int]int) int {
    if len(m) == 0 {
        return 0
    }

    maxKey := -1 << 63
    for k := range m {
        if k > maxKey {
            maxKey = k
        }
    }
    return maxKey
}

func SumOfValuesInMap(m map[int]int) int {
	sum := 0
	for _, v := range m {
		sum += v
	}
	return sum
}


func SwapKeysAndValues(m map[string]string) map[string]string {
	n := map[string]string{}
	for k, v := range m {
		n[v] = k
	}
	return n
}


func CountingSort(contacts []string) map[string]int {
	m := map[string]int{}
	for _, v := range contacts {
		m[v]++
	}
	return m
}



func DeleteLongKeys(m map[string]int) map[string]int {
	for k := range m {
		if len(k) < 6 {
			delete(m, k)
		}
	}
	return m
}

func main() {
	input := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9, "ten": 10}
	fmt.Println(DeleteLongKeys(input))
}