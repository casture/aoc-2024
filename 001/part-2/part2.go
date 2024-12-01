package part_2

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile = "input.txt"
)

func init() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("could not open %s: %v", inputFile, err)
	}
	defer f.Close()

	var left []int
	frequency := make(map[int]int)

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		nums := strings.Split(string(sc.Bytes()), "   ")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		frequency[r] += 1
	}

	sum := 0
	for _, l := range left {
		sum += l * frequency[l]
	}

	log.Printf("Sum: %v", sum)
}
