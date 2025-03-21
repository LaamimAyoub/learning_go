package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("Saying Hello to people", func(t *testing.T) {

		got := Hello("Chriss", "English")
		want := "Hello Chriss"

		assertErrorMessage(t, got, want)

	})

	t.Run("Saying Hello to nothing", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello Nothing"

		assertErrorMessage(t, got, want)

	})

	t.Run("Testing in Spanish", func(t *testing.T) {
		got := Hello("Elodie", spanish)
		want := "Hola Elodie"

		assertErrorMessage(t, got, want)
	})

	t.Run("Testing in French", func(t *testing.T) {
		got := Hello("Elodie", french)
		want := "Bonjour Elodie"

		assertErrorMessage(t, got, want)
	})

	t.Run("Testing in Amazigh", func(t *testing.T) {
		got := Hello("Elodie", "Amazigh")
		want := "Azul Elodie"

		assertErrorMessage(t, got, want)
	})

}

func assertErrorMessage(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

}
