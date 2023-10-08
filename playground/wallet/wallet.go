package wallet

import (
	"errors"
)

type Bitcoin float64

var ErrInvalidAmount = errors.New("invalid amount")

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}


func (w *Wallet) Deposit(amount Bitcoin) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}
	w.balance += amount
	return nil
}

func (w *Wallet) WithDraw(amount Bitcoin) error {
	if amount <= 0 {
		return ErrInvalidAmount
	}

	if amount > w.balance {
		return ErrInvalidAmount
	}

	w.balance -= amount
	return nil
}
