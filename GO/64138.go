package main

import (
    "fmt"
    "unicode"
)

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

func checkPassword(password string) bool {
    return hasMinimumLength(password, 8) && hasUpper(password)
}

func main() {
    // Примеры для тестирования
    fmt.Println(checkPassword("password"))      // false
    fmt.Println(checkPassword("Password123"))   // true
    fmt.Println(checkPassword("Pass"))          // false
    fmt.Println(checkPassword("PASSWORD123"))   // true
    fmt.Println(checkPassword("p@ssw0rd"))     // false
}
