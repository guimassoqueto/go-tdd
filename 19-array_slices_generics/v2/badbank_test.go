package main

import "testing"

func TestBadBank(t *testing.T) {
	transactions := []Transaction {
		{
			From: "Guilherme",
			To: "Karina",
			Sum: 100,
		},
		{
			From: "Guillian",
			To: "Guilherme",
			Sum: 25,
		},
	}
	AssertEqual(t, BalanceFor(transactions, "Karina"), 100)
	AssertEqual(t, BalanceFor(transactions, "Guilherme"), -75)
	AssertEqual(t, BalanceFor(transactions, "Guillian"), -25)
}

func AssertEqual(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}