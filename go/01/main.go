package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/casture/aoc/01/part-1"
	"github.com/casture/aoc/01/part-2"
)

const (
	fileName = "input.txt"
)

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open %s: %v", fileName, err)
	}
	defer f.Close()

	var left, right []int

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		nums := strings.Split(string(sc.Bytes()), "   ")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		left = append(left, l)
		right = append(right, r)
	}

	slices.Sort(left)
	slices.Sort(right)

	d, err := part_1.CalculateDistance(left, right)
	if err != nil {
		log.Fatalf("error while calculating distance: %v", err)
	}
	log.Printf("Distance: %d", d)

	sum, err := part_2.CalculateFrequency(left, right)
	if err != nil {
		log.Fatalf("error while calculating frequency: %v", err)
	}
	log.Printf("Frequency: %d", sum)

}
