package romannumerals

import "testing"

func check_value(t *testing.T, n int, want string) {
	t.Helper()
	got := ConvertToRoman(n)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestRomanNumerals(t *testing.T) {
	check_value(t, 1, "I")
	check_value(t, 2, "II")
	check_value(t, 3, "III")
	check_value(t, 1001, "MI")
	check_value(t, 1003, "MIII")
	check_value(t, 801, "DCCCI")
}
