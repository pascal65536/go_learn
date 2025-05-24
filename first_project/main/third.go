package main

import (
	"fmt"
	"strconv"
)

func SumMix(arr []any) int {
    result := 0
    for _, val := range arr {
        switch v := val.(type) {
        case int:
            result += v
        case string:
            num, err := strconv.Atoi(v)
            if err == nil {
                result += num
            }
        }
    }
    return result
}



// var _ = Describe("Test Example", func() {
//   It("Basic tests", func() {
//     doTest([]any{9, 3, "7", "3"}, 22)
//     doTest([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}, 42)
//     doTest([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2,"0"}, 41)
//     doTest([]any{"1", "5", "8", 8, 9, 9, 2, "3"}, 45)
//     doTest([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}, 61)
//   })
// })


func main() {
	fmt.Println("\nHello, World\n")
	ret := SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7})
	fmt.Println(ret)
}