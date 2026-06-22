package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) error {
	if b.Balance-amount <= 0 {
		return errors.New("Insufficient Funds")
	}
	b.Balance -= amount
	return nil
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("%s: %.2f", b.Owner, b.Balance)
}
func main() {

}
