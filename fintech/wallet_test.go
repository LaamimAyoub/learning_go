package fintech

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, w *Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}

	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("Should have an error here")
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("Check deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(6))

		assertBalance(t, &wallet, Bitcoin(6))

	})

	t.Run("check Withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, &wallet, Bitcoin(5))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(11))
		assertError(t, err, errInsufficientFunds)
		assertBalance(t, &wallet, Bitcoin(10))

	})

}
