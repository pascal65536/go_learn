package main
import "fmt"
import "errors"

var Balance float64 = 0

var (
	ErrAmountIncorrect = errors.New("amount is incorrect")
	ErrBalanceIncorrect = errors.New("balance is incorrect")
)


func topUpBalance(amount float64) error {
	if amount <= 0 {
		return ErrAmountIncorrect
	}
	Balance += amount
	return nil
}

func chargeFromBalance(amount float64) error {
	if amount <= 0 {
		return ErrAmountIncorrect
	}
	if Balance - amount < 0 {
		return ErrBalanceIncorrect
	}
	Balance -= amount
	return nil
}

func TopUpAndGetBalance(amount float64) (float64, error) {
	if err := topUpBalance(amount); err != nil {
		return 0, err
	}
	if Balance < 0 {
		return 0, ErrBalanceIncorrect
	}
	return Balance, nil
}

func ChargeFromAndGetBalance(amount float64) (float64, error) {
	if err := chargeFromBalance(amount); err != nil {
		return 0, err
	}
	if Balance < 0 {
		return 0, ErrBalanceIncorrect
	}
	return Balance, nil
}




func main() {
	fmt.Println(topUpBalance(100000))
	fmt.Println(topUpBalance(150))
	fmt.Println(topUpBalance(150))
	fmt.Println(chargeFromBalance(99999))
	fmt.Println(TopUpAndGetBalance(99999))
	fmt.Println(ChargeFromAndGetBalance(99999))
}

