package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	part_1 "github.com/casture/aoc/02/part-1"
	part_2 "github.com/casture/aoc/02/part-2"
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

	var reports [][]int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		nums := strings.Split(string(sc.Bytes()), " ")
		report := make([]int, len(nums))
		for i, n := range nums {
			report[i], _ = strconv.Atoi(n)
		}
		reports = append(reports, report)
	}

	pt1Safe := part_1.IsSafe(reports)
	log.Printf("Part 1 | Total safe: %d", pt1Safe)

	pt2Safe := part_2.IsSafe(reports)
	log.Printf("Part 2 | Total safe: %d", pt2Safe)
}
