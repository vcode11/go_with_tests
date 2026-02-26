package sum

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers[:])
	want := 15
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumAll(t *testing.T) {
	input_slices := [][]int{
		{1, 2},
		{2, 4},
	}
	got := SumAll(input_slices)
	want := []int{3, 6}
	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)

	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int){
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("Basic test for SumAllTails", func(t *testing.T) {
		got := SumAllTails(
			[]int{1, 2},
			[]int{0, 9},
		)
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("Test array with no tail", func(t *testing.T) {
		got := SumAllTails(
			[]int{1},
			[]int{0, 9},
		)
		want := []int{0, 9}
		checkSums(t, got, want)
	})
	t.Run("Test empty array w", func(t *testing.T) {
		got := SumAllTails(
			[]int{},
			[]int{0, 9},
		)
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
