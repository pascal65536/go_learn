
package main

import "fmt"
import "math"

func main() {
	var a float64
	fmt.Scanln(&a)
	if a < 0 {
		fmt.Println(-1)		
		
	} else {
		resultMessage := fmt.Sprintf("%.3f", math.Sqrt(a))
		fmt.Println(resultMessage)		

	}
}