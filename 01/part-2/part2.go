package part_2

var (
	inputFile = "input.txt"
)

func CalculateFrequency(left, right []int) (int, error) {
	frequency := make(map[int]int)
	for _, r := range right {
		frequency[r] += 1
	}

	sum := 0
	for _, l := range left {
		sum += l * frequency[l]
	}
	return sum, nil
}
