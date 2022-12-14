package main

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"

	"github.com/erikjuhani/advent-of-code-2022/utils"
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

func InputToIntHeap(input []string) *IntHeap {
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

func TopMostCalories(h *IntHeap) int {
	return h.Peek(0)
}

func Top3TotalMostCalories(h *IntHeap) int {
	return h.Peek(0) + h.Peek(1) + h.Peek(2)
}

func Day1() error {
	input, err := utils.ReadInput("./input/day01")
	if err != nil {
		return err
	}

	h := InputToIntHeap(strings.Split(input, "\n"))

	fmt.Printf("DAY 1 (1/2): %d\n", TopMostCalories(h))
	fmt.Printf("DAY 1 (2/2): %d\n", Top3TotalMostCalories(h))

	return nil
}
