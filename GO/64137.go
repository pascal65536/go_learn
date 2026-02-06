package main

import (
    "fmt"
    "math"
)

func findDiscriminant(a, b, c float64) float64 {
    return b*b - 4*a*c
}

func SquareRoots(a, b, c float64) (float64, float64) {
    d := findDiscriminant(a, b, c)
    
    if d < 0 {
        return 0, 0
    }
    
    sqrtD := math.Sqrt(d)
    x1 := (-b - sqrtD) / (2 * a)
    x2 := (-b + sqrtD) / (2 * a)
    
    if d == 0 {
        return x1, x1
    }
    
    if x1 > x2 {
        x1, x2 = x2, x1
    }
    
    return x1, x2
}

func main() {
	fmt.Println(SquareRoots(1, -3, 2))
	fmt.Println(SquareRoots(1, 4, 4.0))
	fmt.Println(SquareRoots(1, 1, 1))
	fmt.Println(SquareRoots(4, 4, 1))
	fmt.Println(SquareRoots(4, 4, -1))
	fmt.Println(SquareRoots(2, 2, 8))
	fmt.Println(SquareRoots(1, 4, -5))
	fmt.Println(SquareRoots(1, 0, -9))
	fmt.Println(SquareRoots(1, 5, 0))
}
