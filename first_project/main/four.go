package main

import (
    "unicode"
    "strings"
)

func Accum(s string) string {
    var total []string
    for i, r := range []rune(s) {
        var letters []rune
        f := true
        for j := 0; j <= i; j++ {
            if f {
                letters = append(letters, unicode.ToUpper(r))
            } else {
                letters = append(letters, unicode.ToLower(r))
            }
            f = false
        }        
        total = append(total, string(letters))
    }    
    return strings.Join(total, "-")
}


func main() {
    ret := Accum("HbideVbxncC")
    fmt.Println(ret)
}


// func Digits(n uint64) int {
//     if n == 0 {
//         return 1
//     }
//     digits := 0
//     for n > 0 {
//       digits++
//       n /= 10
//     }
//     return digits
// }

// func Digitize(n int) []int {
//     var digits []int
//     if n == 0 {
//       return []int{0}
//     }

//     for n > 0 {
//       digits = append(digits, n%10)
//       n /= 10
//     }
//     return digits
//   }


// var _ = Describe("Test Example", func() {
// It("Basic tests", func() {
// doTest([]any{9, 3, "7", "3"}, 22)
// doTest([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7}, 42)
// doTest([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2,"0"}, 41)
// doTest([]any{"1", "5", "8", 8, 9, 9, 2, "3"}, 45)
// doTest([]any{8, 0, 0, 8, 5, 7, 2, 3, 7, 8, 6, 7}, 61)
// ret := SumMix([]any{"5", "0", 9, 3, 2, 1, "9", 6, 7})
// fmt.Println(ret)


