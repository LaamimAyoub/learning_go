package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Run with Array", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

	t.Run("Run with slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("Test Sum all", func(t *testing.T) {
		arr1 := []int{1, 2, 3}
		arr2 := []int{4, 5, 6}

		got := SumAll(arr1, arr2)
		want := []int{6, 15}

		if !slices.Equal(got, want) {
			t.Error("Sum all failed")
		}

	})
}

func TestSumAllTail(t *testing.T) {
	t.Run("Test sum tail", func(t *testing.T) {
		arr1 := []int{1, 2, 3}
		arr2 := []int{4, 5, 6}

		got := SumAllTail(arr1, arr2)
		want := []int{5, 11}

		if !slices.Equal(got, want) {
			t.Error("Sum all failed")
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
