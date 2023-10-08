package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("wallet: should return a balance of 0 when balance init", func(t *testing.T) {
		w := Wallet{}
		got := w.Balance()
		want := Bitcoin(0)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("wallet: should return an error if zero is deposited", func(t *testing.T) {
		w := Wallet{}
		got := w.Deposit(0)
		want := ErrInvalidAmount

		assertEqual(t, got, want)
	})

	t.Run("wallet: should return an error if negative is deposited", func(t *testing.T) {
		w := Wallet{}
		got := w.Deposit(-5.5)
		want := ErrInvalidAmount

		assertEqual(t, got, want)
	})

	t.Run("wallet: should a correct value after an successful deposit", func(t *testing.T) {
		w := Wallet{}
		err := w.Deposit(20)
		got := w.Balance()
		want := Bitcoin(20)

		assertEqual(t, err, nil)
		assertEqual(t, got, want)
	})

	t.Run("wallet: should return an error if amount is negative", func(t *testing.T) {
		w := Wallet{}
		got := w.WithDraw(-10)
		want := ErrInvalidAmount

		assertEqual(t, got, want)
	})

	t.Run("wallet: should return an error if amount is greater than the balance", func(t *testing.T) {
		w := Wallet{}
		got := w.WithDraw(20)
		want := ErrInvalidAmount

		assertEqual(t, got, want)
	})

	t.Run("wallet: should the correct amount after a succeful withdraw", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(500)
		w.WithDraw(20)
		got := w.Balance()
		want := Bitcoin(480)

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertEqual(t testing.TB, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v got %v", got, want)
	}
}