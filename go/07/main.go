package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	debug = true
)

func main() {
	var equations []Equation

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := strings.Split(sc.Text(), ":")
		sol, _ := strconv.Atoi(line[0])
		ops := nums(line[1])
		equations = append(equations, Equation{
			solution: sol,
			operands: ops,
		})
	}

	sum := 0
	for i, e := range equations {
		if e.isValid() {
			log.Println("equation is valid: ", i)
			sum += e.solution
		}
	}
	log.Println(sum)
}

type Equation struct {
	solution int
	operands []int
}

func (e *Equation) isValid() bool {
	c := getCombinations(len(e.operands) - 1)
	for _, ops := range c {
		agg := e.operands[0]
		for j, o := range ops {
			switch o {
			case ADD:
				agg += e.operands[j+1]
			case MULT:
				agg *= e.operands[j+1]
			}
		}
		if agg == e.solution {
			return true
		}
	}
	return false
}

type Operator int

const (
	ADD = iota
	MULT
	SUB
	DIV
)

func getCombinations(max int) [][]Operator {
	var c [][]Operator
	c = append(c, []Operator{ADD})
	c = append(c, []Operator{MULT})
	for i := 0; i < max-1; i++ {
		length := len(c)
		for j := 0; j < length; j++ {
			c = append(c, append(deepClone(c[j]), MULT))
			c[j] = append(c[j], ADD)
		}
	}
	return c
}

func deepClone(slice []Operator) []Operator {
	clone := make([]Operator, len(slice))
	for i, v := range slice {
		clone[i] = v
	}
	return clone
}

// parse strings separated by spaces into an []int
func nums(s string) (n []int) {
	for _, v := range strings.Fields(s) {
		i, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		n = append(n, i)
	}
	return
}
