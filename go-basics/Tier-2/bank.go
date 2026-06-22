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
	if b.Balance < amount {
		return errors.New("Insufficient Funds")
	}
	b.Balance -= amount
	return nil //errors.New("Withdrawal Successful")
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("%s: $%.2f", b.Owner, b.Balance)
}
func TaskThree() {
	account := BankAccount{Owner: "Offorbuzor Divine", Balance: 225000.76}
	account.Deposit(7500)
	err := account.Withdraw(5000)
	fmt.Println(err)
	err = account.Withdraw(50000000)
	fmt.Println(err)
	statement := account.String()
	fmt.Println(statement)
}
