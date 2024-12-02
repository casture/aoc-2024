package part_1

func CalculateDistance(left, right []int) (int, error) {
	var dif int
	for i, l := range left {
		dif += abs(l - right[i])
	}
	return dif, nil
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
