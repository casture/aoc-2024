package main

import (
	"log"
	"os"
	"strings"
)

var (
	fileName = "input.txt"
	debug    = false
)

func debugLog(m string, args ...any) {
	if debug {
		log.Printf("[Debug] "+m, args...)
	}
}

func main() {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Could not open file=%s: %v", fileName, err)
	}

	lines := strings.Split(strings.TrimRight(string(f), "\n"), "\n")
	grid := make([][]rune, len(lines))
	for i, l := range lines {
		grid[i] = []rune(l)
	}

	count := getWordCount("XMAS", grid)
	log.Printf("Word count: %d\n", count)

	count = getXCount(grid)
	log.Printf("X pattern count: %d\n", count)
}

func getXCount(grid [][]rune) int {
	count := 0
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if grid[i][j] == 'A' {
				if (grid[i-1][j-1] == 'M' && grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'M' && grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' && grid[i+1][j+1] == 'S') || (grid[i-1][j-1] == 'S' && grid[i-1][j+1] == 'S' && grid[i+1][j-1] == 'M' && grid[i+1][j+1] == 'M') || (grid[i-1][j-1] == 'S' && grid[i-1][j+1] == 'M' && grid[i+1][j-1] == 'S' && grid[i+1][j+1] == 'M') {
					count++
				}
			}
		}
	}
	return count
}

func getWordCount(word string, grid [][]rune) int {
	count := 0
	for i, row := range grid {
		for j := range row {

			canCheck := func(direction direction) bool {
				switch direction {
				case UP:
					return i >= len(word)-1
				case LEFT:
					return j >= len(word)-1
				case DOWN:
					return i <= len(grid)-len(word)
				case RIGHT:
					return j <= len(row)-len(word)
				}
				return false
			}

			var checkers []checker
			if canCheck(UP) {
				checkers = append(checkers, checker{delta{0, -1}, i, j})
				if canCheck(RIGHT) {
					checkers = append(checkers, checker{delta{1, -1}, i, j})
				}
				if canCheck(LEFT) {
					checkers = append(checkers, checker{delta{-1, -1}, i, j})
				}
			}
			if canCheck(DOWN) {
				checkers = append(checkers, checker{delta{0, 1}, i, j})
				if canCheck(RIGHT) {
					checkers = append(checkers, checker{delta{1, 1}, i, j})
				}
				if canCheck(LEFT) {
					checkers = append(checkers, checker{delta{-1, 1}, i, j})
				}
			}
			if canCheck(RIGHT) {
				checkers = append(checkers, checker{delta{1, 0}, i, j})
			}
			if canCheck(LEFT) {
				checkers = append(checkers, checker{delta{-1, 0}, i, j})
			}

			debugLog("%+v\n", checkers)

			for _, checker := range checkers {
				if checker.check([]rune(word), grid) {
					count++
					debugLog("count=%d dir=%v\n", count, checker.delta)
				}
			}
		}
	}
	return count
}

func (c *checker) check(word []rune, grid [][]rune) bool {
	debugLog("%+v", string(word))
	for _, char := range word {
		debugLog("%+v", c)
		debugLog("%c == %c", char, grid[c.i][c.j])
		if char != grid[c.i][c.j] {
			return false
		}
		c.move()
	}
	return true
}

type checker struct {
	delta
	i, j int
}

type direction int

func (c *checker) move() {
	c.i += c.y
	c.j += c.x
}

const (
	UP direction = iota
	RIGHT
	DOWN
	LEFT
)

type delta struct {
	x, y int
}
