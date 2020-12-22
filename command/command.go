package main

import "fmt"

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposit", amount, "\b,balance is now", b.balance)
}

func main() {

}
