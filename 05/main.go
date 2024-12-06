package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
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

const (
	RULES = iota
	PAGES
)

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("could not open file=%s: %+v", fileName, err)
	}
	defer f.Close()

	rules := make(map[string]map[string]bool)
	var updates [][]string

	mode := RULES
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			mode = PAGES
			continue
		}
		if mode == RULES {
			rel := strings.Split(line, "|")
			p, c := rel[0], rel[1]
			if rules[p] == nil {
				rules[p] = make(map[string]bool)
			}
			rules[p][c] = true
		} else if mode == PAGES {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	debugLog("", updates)

	var correctUpdates [][]string

outer:
	for _, u := range updates {
		mustFollow := make(map[string]bool)
		for i := len(u) - 1; i >= 0; i-- {
			if _, ok := mustFollow[u[i]]; ok {
				continue outer
			}
			for key := range rules[u[i]] {
				mustFollow[key] = true
			}
			debugLog("%+v", mustFollow)
		}
		correctUpdates = append(correctUpdates, u)
	}

	debugLog("num correctUpdates: %d\n", len(correctUpdates))

	sum := 0
	for _, u := range correctUpdates {
		n, _ := strconv.Atoi(u[(len(u)-1)/2])
		sum += n
	}

	log.Printf("sum=%d\n", sum)
}
