package main

import "fmt"

type Wallet struct {
	balance float64
}

func (w *Wallet) Balance() float64 {
	return w.balance
}

func (w *Wallet) Deposit(amount float64) float64 {
	fmt.Printf("address of balance in test is %v \n", &w.balance)
	w.balance += amount
	return 10
}