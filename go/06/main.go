package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	fileName = "input.txt"
	debug    = false
)

func debugLog(m string, args ...any) {
	if debug {
		log.Printf("[Debug] "+m, args...)
	}
}

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open file=%s: %+v", fileName, err)
	}
	defer f.Close()

	obstacles := make(map[string]bool)
	var start position
	var orientation Orientation

	sc := bufio.NewScanner(f)
	maxX := 0
	i := 0
	for sc.Scan() {
		line := []rune(sc.Text())
		maxX = len(line)
		for j, r := range line {
			if r == '.' {
				continue
			}
			p := position{j, i}
			if r == '#' {
				obstacles[p.string()] = true
				continue
			}
			start = p
			if r == '^' {
				orientation = UP
			} else if r == '>' {
				orientation = RIGHT
			} else if r == 'v' {
				orientation = DOWN
			} else {
				orientation = LEFT
			}
		}
		i++
	}
	g := &guard{
		start,
		orientation,
		make(map[string]bool),
		chart{
			obstacles,
			maxX,
			i,
		},
		make(map[string][]Orientation),
	}

	debugLog("%+v", g)

	//PART 1

	// add starting position to visited
	pKey := position{g.x, g.y}.string()
	g.visited[pKey] = true

loop:
	for {
		switch g.peek() {
		case OK:
			g.move()
			pKey := position{g.x, g.y}.string()
			g.visited[pKey] = true
		case OBSTACLE:
			g.turn()
		case OOB:
			break loop
		}
	}

	log.Printf("total visited: %d", g.routeLength())

	// PART 2
	totalLoops := 0
	for key := range g.visited {
		p := newPosition(key)
		if start.x == p.x && start.y == p.y {
			continue
		}

		debugLog("new o: %s", key)

		// clone starting obstacles
		dObstacles := make(map[string]bool)
		for k, v := range obstacles {
			dObstacles[k] = v
		}
		// add new obstacle
		dObstacles[key] = true

		// create new guard
		g := &guard{
			start,
			orientation,
			make(map[string]bool),
			chart{
				dObstacles,
				maxX,
				i,
			},
			make(map[string][]Orientation),
		}

	loop2:
		for {
			switch g.peek() {
			case OK:
				g.move()
				pKey := position{g.x, g.y}.string()
				g.visited[pKey] = true
			case OBSTACLE:
				g.turn()
			case OOB:
				break loop2
			case LOOP:
				debugLog("loop detected: %s", g.Orientation.string())
				totalLoops++
				break loop2
			}
		}
	}

	log.Printf("total loops: %d", totalLoops)
}

type guard struct {
	position
	Orientation
	visited       map[string]bool
	chart         chart
	seenObstacles map[string][]Orientation
}

type chart struct {
	obstacles  map[string]bool
	maxX, maxY int
}

func (c chart) outOfBounds(x, y int) bool {
	return x < 0 || y < 0 || x > c.maxX || y > c.maxY
}

func (c chart) isObstacle(x, y int) bool {
	_, ok := c.obstacles[position{x, y}.string()]
	return ok
}

type peekResult int

const (
	OK peekResult = iota
	OOB
	OBSTACLE
	LOOP
)

func (g guard) peek() peekResult {
	x, y := g.getNewPosition()
	if g.chart.outOfBounds(x, y) {
		debugLog("%d-%d OOB", y, x)
		return OOB
	} else if g.chart.isObstacle(x, y) {
		pKey := position{x, y}.string()
		for _, o := range g.seenObstacles[pKey] {
			if o == g.Orientation {
				return LOOP
			}
		}
		g.seenObstacles[pKey] = append(g.seenObstacles[pKey], g.Orientation)
		debugLog("%d-%d OBSTACLE", y, x)
		return OBSTACLE
	} else {
		debugLog("%d-%d OK", y, x)
		return OK
	}
}

func (g *guard) move() {
	g.x, g.y = g.getNewPosition()
}

func (g guard) getNewPosition() (dX, dY int) {
	switch g.Orientation {
	case UP:
		dX, dY = g.x, g.y-1
	case DOWN:
		dX, dY = g.x, g.y+1
	case LEFT:
		dX, dY = g.x-1, g.y
	case RIGHT:
		dX, dY = g.x+1, g.y
	}
	return
}

func (g guard) routeLength() int {
	return len(g.visited)
}

func (g *guard) turn() {
	g.Orientation = (g.Orientation + 1) % 4
	debugLog("turned %s", g.Orientation.string())
}

type position struct {
	x, y int
}

func newPosition(s string) position {
	p := strings.Split(s, "-")
	x, _ := strconv.Atoi(p[1])
	y, _ := strconv.Atoi(p[0])
	return position{y, x}
}

func (p position) string() string {
	return fmt.Sprintf("%d-%d", p.y, p.x)
}

type Orientation int

func (o Orientation) string() string {
	switch o {
	case UP:
		return "UP"
	case RIGHT:
		return "RIGHT"
	case DOWN:
		return "DOWN"
	case LEFT:
		return "LEFT"
	}
	return "Unknown"
}

const (
	UP Orientation = iota
	RIGHT
	DOWN
	LEFT
)
