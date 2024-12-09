package part_1

func IsSafe(reports [][]int) int {
	total := 0
	for _, report := range reports {
		if IsReportSafe(report) {
			total += 1
		}
	}
	return total
}

func IsReportSafe(report []int) bool {
	direction := 1
	for i := 0; i < len(report)-1; i++ {
		if i == 0 && report[i] > report[i+1] {
			direction = -1
		}
		dif := report[i] - report[i+1]
		if abs(dif) < 1 || abs(dif) > 3 || dif*direction > 0 {
			return false
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
