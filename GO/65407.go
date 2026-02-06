package main

import (
	"fmt"

	"config/config"
)

func main() {
	cfg := config.NewConfig()
	fmt.Printf("Баланс по умолчанию: %.2f\n", cfg.DefaultBalance)
}
