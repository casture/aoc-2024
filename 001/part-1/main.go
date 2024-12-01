package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	inputFile = "input.txt"
)

func main() {
	f, err := os.Open(inputFile)
	if err != nil {

		log.Fatalf("could not open %s: %v", inputFile, err)
	}
	defer f.Close()

	var lA, rA []int
	var lines int

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines += 1
		nums := strings.Split(string(sc.Bytes()), "   ")
		l, _ := strconv.Atoi(nums[0])
		r, _ := strconv.Atoi(nums[1])
		lA = append(lA, l)
		rA = append(rA, r)
	}

	QuickSort(lA)
	QuickSort(rA)

	var dif float64
	for i, li := range lA {
		dif += math.Abs(float64(li) - float64(rA[i]))
	}
	log.Printf("Dif: %v", dif)
}

// QuickSort sorts an array using the Quick Sort algorithm
func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pivotIndex := partition(arr, low, high)
		quickSortHelper(arr, low, pivotIndex-1)  // Sort left of pivot
		quickSortHelper(arr, pivotIndex+1, high) // Sort right of pivot
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // Swap smaller element with i-th position
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1] // Place pivot in correct position
	return i + 1
}
