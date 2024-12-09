package part_2

import "testing"

func TestIsSafe(t *testing.T) {
	t.Run("example input", func(t *testing.T) {
		reports := [][]int{
			{7, 6, 4, 2, 1},
			{1, 2, 7, 8, 9},
			{9, 7, 6, 2, 1},
			{1, 3, 2, 4, 5},
			{8, 6, 4, 4, 1},
			{1, 3, 6, 7, 9},
		}
		r := IsSafe(reports)
		if r != 4 {
			t.Fail()
		}
	})
	t.Run("l should skip unsafe index", func(t *testing.T) {
		reports := [][]int{
			{17, 19, 17, 20, 23},
		}
		r := IsSafe(reports)
		if r != 1 {
			t.Fail()
		}
	})
	t.Run("should be safe without zero index", func(t *testing.T) {
		reports := [][]int{
			{20, 17, 19, 20, 23},
		}
		r := IsSafe(reports)
		if r != 1 {
			t.Fail()
		}
	})
}
