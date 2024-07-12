package algos

import (
	"reflect"
	"testing"
)

// test cases
var testCases = []struct {
	name  string
	input []int
	want  []int
}{
	{
		name:  "already sorted",
		input: []int{1, 2, 3, 4, 5},
		want:  []int{1, 2, 3, 4, 5},
	},
	{
		name:  "reverse sorted",
		input: []int{5, 4, 3, 2, 1},
		want:  []int{1, 2, 3, 4, 5}},
	// with negative numbers
	{
		name:  "negative numbers",
		input: []int{1, 2, 3, 4, 5, -1, -2, -3, -4, -5},
		want:  []int{-5, -4, -3, -2, -1, 1, 2, 3, 4, 5}},
	{
		name:  "negative numbers reverse",
		input: []int{5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5},
		want:  []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5},
	},
	// edge cases
	{
		name:  "empty slice",
		input: []int{},
		want:  []int{}},
	{
		name:  "single element",
		input: []int{1},
		want:  []int{1},
	},
	{
		name:  "all same elements",
		input: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		want:  []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	},
	// cases where naive quicksort is inefficient
	{
		name:  "already sorted",
		input: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		name:  "reverse sorted",
		input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		want:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	// worst case scenario
	{
		name:  "worst case scenario",
		input: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1, -2, -3, -4, -5},
		want:  []int{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
}

func TestQuickSort(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := QuickSort(tc.input, 0, len(tc.input)-1)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
