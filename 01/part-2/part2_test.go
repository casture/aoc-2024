package part_2

import (
	"testing"
)

func TestCalculateDistance(t *testing.T) {
	left := []int{3, 4, 2, 1, 3, 3}
	right := []int{4, 3, 5, 3, 9, 3}

	d, err := CalculateFrequency(left, right)

	if err != nil {
		t.FailNow()
	}
	if d != 31 {
		t.FailNow()
	}
}
