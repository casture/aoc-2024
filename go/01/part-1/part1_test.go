package part_1

import (
	"slices"
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}
	slices.Sort(left)
	slices.Sort(right)

	d, err := CalculateDistance(left, right)

	if err != nil {
		t.FailNow()
	}
	if d != 11 {
		t.FailNow()
	}
}
