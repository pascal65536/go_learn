package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
	owner   string
}

func NewAccount(owner string) *Account {
	return &Account{
		owner:   owner,
		balance: 0.0,
	}
}

func (a *Account) SetBalance(value float64) error {
	if value < 0 {
		return errors.New("баланс не может быть отрицательным")
	}
	a.balance = value
	return nil
}

func (a *Account) GetBalance() float64 {
	return a.balance
}

func (a *Account) Deposit(value float64) error {
	if value <= 0 {
		return errors.New("сумма пополнения должна быть положительной")
	}
	a.balance += value
	return nil
}

func (a *Account) Withdraw(value float64) error {
	if value <= 0 {
		return errors.New("сумма снятия должна быть положительной")
	}
	if a.balance < value {
		return errors.New("недостаточно средств на счете")
	}
	a.balance -= value
	return nil
}

func main() {
	acc := NewAccount("Арсений")
	
	// Пример использования
	acc.Deposit(1000)
	fmt.Printf("Баланс после пополнения: %.2f\n", acc.GetBalance()) // 1000.00
	
	err := acc.Withdraw(300)
	if err != nil {
		fmt.Println("Ошибка снятия:", err)
	} else {
		fmt.Printf("Баланс после снятия: %.2f\n", acc.GetBalance()) // 700.00
	}
	
	err = acc.SetBalance(-50)
	if err != nil {
		fmt.Println("Ошибка установки баланса:", err) // Баланс не может быть отрицательным
	}
}
