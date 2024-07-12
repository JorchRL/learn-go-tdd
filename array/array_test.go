package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Sum of an arbitrary collection of numbers", func(t *testing.T) {
	nums := []int{1,2,3}
	got := Sum(nums)
	want := 6

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}	})
}

func TestSumAll(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}}

	t.Run("Sum of 2 slices", func(t *testing.T) {
		got := SumAll([]int{1,2}, []int{0,9})
		want := []int{3, 9}
		
		checkSums(t, got, want)
	})

	t.Run ("Sum of 3 slices", func(t *testing.T) {
		got := SumAll([]int{1,2}, []int{0,9}, []int{1,1})
		want := []int{3, 9, 2}
		
		checkSums(t, got, want)
	})

	t.Run("Sum when one of the slices is empty", func(t *testing.T) {
		got := SumAll([]int{}, []int{3,4,5})
		want := []int{0, 12}

		checkSums(t, got, want)
	})
}

func TestSumAllTais(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("Sum of tails of 2 slices", func(t *testing.T) {
		got := SumAllTails([]int{1,2,3}, []int{4,3,2})
		want := []int{5, 5}
		checkSums(t, got, want)
	})

	t.Run("SumAllTails with an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}