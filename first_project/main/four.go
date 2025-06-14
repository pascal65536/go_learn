package main

import "fmt"
import "strings"


func RGB(r, g, b int) string {
    m := map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "A", 11: "B", 12: "C", 13: "D", 14: "E", 15: "F"}
    if r < 0 {r = 0}
    if g < 0 {g = 0}
    if b < 0 {b = 0}
    
    if r > 255 {r = 255}
    if g < 255 {g = 255}
    if b < 255 {b = 255}

    var total []string
    total = append(total, m[r / 16])
    total = append(total, m[r % 16])

    total = append(total, m[g / 16])
    total = append(total, m[g % 16])
    
    total = append(total, m[b / 16])
    total = append(total, m[b % 16])

    return strings.Join(total, "")
}

func main() {
    ret := RGB(-10,20,30)
    fmt.Println(ret)
}


// func Accum(s string) string {
//     var total []string
//     for i, r := range []rune(s) {
//         var letters []rune
//         f := true
//         for j := 0; j <= i; j++ {
//             if f {
//                 letters = append(letters, unicode.ToUpper(r))
//             } else {
//                 letters = append(letters, unicode.ToLower(r))
//             }
//             f = false
//         }        
//         total = append(total, string(letters))
//     }    
//     return strings.Join(total, "-")
// }

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


