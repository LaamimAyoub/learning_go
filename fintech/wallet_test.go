package fintech

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(6))

	got := wallet.Balance()
	want := Bitcoin(6)

	if got != want {
		t.Errorf("got %2d, want %2d", got, want)
	}

}
