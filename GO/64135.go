package main

import "fmt"


func BuyFries(size string) {
	if size == "S" {
		printPrice(49, "Картошка фри")
	} else if size == "M" {
		printPrice(89, "Картошка фри")
	} else if size == "L" {
		printPrice(99, "Картошка фри")	
	} else {
		fmt.Println("Некорректный размер")	
	}
}


func BuyCola(size string) {
	if size == "S" {
		printPrice(99, "Кола")
	} else if size == "M" {
		printPrice(119, "Кола")
	} else if size == "L" {
		printPrice(139, "Кола")	
	} else {
		fmt.Println("Некорректный размер")	
	}
}


func printPrice(cost int, product string) {
	result := fmt.Sprintf("%s будет стоить %d рублей", product, cost)
	fmt.Println(result)	
}



func main() {
	BuyFries("L")
	BuyFries("M")
	BuyFries("S")

	BuyCola("S")
	BuyCola("M")
	BuyCola("L")
}