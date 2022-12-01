package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

// This list represents the Calories of the food carried by five Elves:
//
//    The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
//    The second Elf is carrying one food item with 4000 Calories.
//    The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
//    The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
//    The fifth Elf is carrying one food item with 10000 Calories.
//
// In case the Elves get hungry and need extra snacks, they need to know which
// Elf to ask: they'd like to know how many Calories are being carried by the
// Elf carrying the most Calories. In the example above, this is 24000 (carried
// by the fourth Elf).
//
// Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?

func readInput() ([]string, error) {
	inputFile, err := os.Open("./day1.input")
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func inputToIntHeap(input []string) *IntHeap {
	h := &IntHeap{}
	heap.Init(h)

	var sum int
	for _, line := range input {
		n, err := strconv.Atoi(line)
		if err != nil {
			heap.Push(h, sum)
			sum = 0
			continue
		}
		sum += n
	}

	return h
}

func topMostCalories(h *IntHeap) int {
	return h.Peek(0)
}

func top3TotalMostCalories(h *IntHeap) int {
	return h.Peek(0) + h.Peek(1) + h.Peek(2)
}

func day1() error {
	input, err := readInput()
	if err != nil {
		return err
	}

	h := inputToIntHeap(input)

	fmt.Printf("DAY 1 (1/2): %d\n", topMostCalories(h))
	fmt.Printf("DAY 2 (2/2): %d\n", top3TotalMostCalories(h))
	return nil
}
