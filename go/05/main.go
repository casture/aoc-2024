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
	debug    = true
)

func debugLog(m string, args ...any) {
	if debug {
		log.Printf("[Debug] "+m, args...)
	}
}

type mode int

const (
	RULES mode = iota
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

	var correctUpdates, incorrectUpdates [][]string

	for _, update := range updates {
		mustFollow := make(map[string]int)
		correct := true
		for i := len(update) - 1; i >= 0; i-- {
			page := update[i]
			updateIndex := i
			if toIndex, ok := mustFollow[page]; ok {
				correct = false
				debugLog("---%+v", update)
				movePageToIndex(update, i, toIndex)
				updateIndex = toIndex
				// update mustFollow indices for all affected pages
				for k, v := range mustFollow {
					if i <= v && v <= toIndex {
						mustFollow[k] = v - 1
					}
				}
				debugLog("+++%+v\n\n", update)
			}
			for key := range rules[page] {
				if _, ok := mustFollow[key]; !ok {
					mustFollow[key] = updateIndex
				}
			}
			// debugLog("%+v", mustFollow)
		}
		if correct {
			correctUpdates = append(correctUpdates, update)
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}
	}

	debugLog("num correctUpdates: %d\n", len(correctUpdates))
	debugLog("num incorrectUpdates: %d\n", len(incorrectUpdates))

	sum := 0
	for _, u := range correctUpdates {
		n, _ := strconv.Atoi(u[(len(u)-1)/2])
		sum += n
	}
	log.Printf("correct updates sum=%d\n", sum)

	sum = 0
	for _, u := range incorrectUpdates {
		n, _ := strconv.Atoi(u[(len(u)-1)/2])
		sum += n
	}
	log.Printf("fixed updates sum=%d\n", sum)

	list := []string{"a", "b", "c"}
	movePageToIndex(list, 0, 2)
	debugLog("%+v", list)

}

func movePageToIndex(row []string, from, to int) {
	direction := 1
	if from-to > 0 {
		direction = -1
	}
	for i, j := from, from+direction; i < to; i, j = i+direction, j+direction {
		row[i], row[j] = row[j], row[i]
	}
}
