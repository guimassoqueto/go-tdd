package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)
	want := 10.0

	if got != want {
		t.Errorf("got %f want %f", got, want)
	}
}