package main

import "unicode"
    

func hasMinimumLength(password string, minLength int) bool {
    return len(password) >= minLength
}

func hasUpper(password string) bool {
    for _, char := range password {
        if unicode.IsUpper(char) {
            return true
        }
    }
    return false
}

func hasLowerCase(password string) bool {
    for _, char := range password {
        if unicode.IsLower(char) {
            return true
        }
    }
    return false
}

func ratePassword(password string) string {
    checksPassed := 0
    
    if hasMinimumLength(password, 8) {
        checksPassed++
    }
    if hasUpper(password) {
        checksPassed++
    }
    if hasLowerCase(password) {
        checksPassed++
    }
    
    switch checksPassed {
    case 0:
        return "Пароль недостаточно безопасен, придумайте новый"
    case 1:
        return "Слабый пароль"
    case 2:
        return "Средний пароль"
    case 3:
        return "Сложный пароль"
    }
    return ""
}

import "fmt"

func main() {
	fmt.Println(ratePassword("abcdefGhjhg"))
	fmt.Println(ratePassword("aaaaA"))
	fmt.Println(ratePassword("ffffffffffffffff"))
	fmt.Println(ratePassword("OMGGGGGGGGGGGG"))
	fmt.Println(ratePassword("dfghTnbH"))
	fmt.Println(ratePassword(""))
	fmt.Println(ratePassword("Y"))
}