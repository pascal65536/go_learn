package main

import "fmt"


func PrettyArrayOutput(array [9]string) {
    for i := 0; i < len(array); i++ {
        if i < 7 {
            fmt.Printf(fmt.Sprintf("%d я уже сделал: %s\n", i+1, array[i]))
        } else {
            fmt.Printf(fmt.Sprintf("%d не успел сделать: %s\n", i+1, array[i]))
        }
    }
}

func SumOfArray(array [6]int) int {
    sum := 0
    for i := 0; i < len(array); i++ {
        sum += array[i]
    }
    return sum
}

func FindMinMaxInArray(array [10]int) (int, int) {
    min := array[0]
    max := array[0]
    for i := 1; i < len(array); i++ {
        if min > array[i] {
            min = array[i]
        }
        if max < array[i]  {
            max = array[i]
        }
    }
    return min, max
}


func FiveSteps(array [5]int) [5]int {
    new := [5]int{array[4], array[3], array[2], array[1], array[0]}
    return new
}

func ThirdElementInArray(array [6]int) int {
    return array[2]
}



func main() {
	input := [9]string{
		"проснуться",
		"позавтракать",
		"сходить в школу",
		"пообедать",
		"погулять с друзьями",
		"сделать домашнюю работу",
		"попрограммировать на Go",
		"поужинать",
		"лечь спать",
	}
	PrettyArrayOutput(input)
}