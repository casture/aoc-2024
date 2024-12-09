package part_2

import (
	part_1 "github.com/casture/aoc/02/part-1"
)

func IsSafe(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if part_1.IsReportSafe(report) {
			total++
			continue
		}

		for i := range report {
			newReport := removeIndex(report[:], i)
			if part_1.IsReportSafe(newReport) {
				total++
				break
			}
		}
	}
	return total
}

func removeIndex(s []int, index int) []int {
	if index < 0 || index >= len(s) {
		return s
	}
	newSlice := make([]int, 0, len(s)-1)
	newSlice = append(newSlice, s[:index]...)
	newSlice = append(newSlice, s[index+1:]...)
	return newSlice
}

func isUnsafe(dif, direction int) bool {
	return abs(dif) < 1 || abs(dif) > 3 || dif*direction > 0
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
