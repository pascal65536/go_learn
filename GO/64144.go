package main

import "fmt"

func PrintComplexNumber(z complex64) {
	realPart := real(z)
	imagPart := imag(z)
	fmt.Printf("Действительная часть: %.2f. Мнимая часть: %.2f.\n", realPart, imagPart)
}

func main() {
	var z complex64 = complex(3.14159, -2.71828)
	PrintComplexNumber(z)
	z = complex(-0.0049, 0.0051)
	PrintComplexNumber(z)
}