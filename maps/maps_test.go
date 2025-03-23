package MyMaps

import (
	"testing"
)

func TestMaps(t *testing.T) {

	t.Run("Look for word", func(t *testing.T) {
		d := dictionary{"test": "this is just a test"}

		got, err := d.Search("test")

		if err != nil {
			t.Fatal("Should not fail here")
		}

		want := "this is just a test"

		assertStrings(t, got, want)

	})

	t.Run("Uknown word", func(t *testing.T) {
		d := dictionary{"test": "this is just a test"}

		_, err := d.Search("Unkown")
		assertError(t, err, errNotFound)

	})

}

func TestAdd(t *testing.T) {

	t.Run("Add a new word", func(t *testing.T) {
		d := dictionary{}
		err := d.Add("test", "this is just a test")

		assertError(t, err, nil)

		want := "this is just a test"
		got, err := d.Search("test")

		assertError(t, err, nil)

		assertStrings(t, got, want)

	})

	t.Run("Add an existing word", func(t *testing.T) {
		d := dictionary{"test": "test"}
		err := d.Add("test", "new test")

		assertError(t, err, errAlreadyExists)

	})

}

func assertStrings(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
