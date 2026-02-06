package main

import "unicode"
import "fmt"


func CheckOnlyASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return false
		}
	}
	return true
}



func main() {
	fmt.Println(CheckOnlyASCII("Привет,"))
	fmt.Println(CheckOnlyASCII("不好意思"))
}

