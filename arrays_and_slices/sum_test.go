package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		want := 6
		got := Sum(numbers)

		if got != want {
			t.Errorf("got %d, wanted %d, given %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	firstSlice := []int{5, 4, 1}
	secondSlice := []int{7, 6, 5, 2}

	want := []int{10, 20}
	got := SumAll(firstSlice, secondSlice)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	// TODO: this
}
