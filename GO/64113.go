
package main

import "fmt"
import "math"

func main() {
	var raz float64
	fmt.Scanln(&raz)
	var dva float64
	fmt.Scanln(&dva)

	if raz > dva {
		fmt.Println(math.Round(raz))
	} else {
		fmt.Println(math.Round(dva))
	}
}