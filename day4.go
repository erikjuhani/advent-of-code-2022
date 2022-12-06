package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/erikjuhani/advent-of-code-2022/utils"
)

type Range struct {
	start, end int
}

func (r Range) Contains(rb Range) bool {
	if r.start >= rb.start && r.end <= rb.end {
		return true
	}

	return false
}

func (r Range) ContainsAny(rb Range) bool {
	if r.end >= rb.start && r.start <= rb.end {
		return true
	}

	return false
}

func NewRange(rawRange []string) (Range, error) {
	start, err := strconv.Atoi(rawRange[0])
	if err != nil {
		return Range{}, err
	}

	end, err := strconv.Atoi(rawRange[1])
	if err != nil {
		return Range{}, err
	}

	return Range{start, end}, nil
}

func Day4() error {
	input, err := utils.ReadInput("./input/day04")
	if err != nil {
		return err
	}

	var amount1, amount2 int
	for _, s := range strings.Split(input, "\n") {
		pair := strings.Split(s, ",")
		if len(pair) != 2 {
			continue
		}
		rangeA, err := NewRange(strings.Split(pair[0], "-"))
		if err != nil {
			return err
		}
		rangeB, err := NewRange(strings.Split(pair[1], "-"))
		if err != nil {
			return err
		}

		if rangeA.Contains(rangeB) || rangeB.Contains(rangeA) {
			amount1++
		}

		if rangeA.ContainsAny(rangeB) || rangeB.ContainsAny(rangeA) {
			amount2++
		}
	}

	fmt.Printf("DAY 4 (1/2): %d\n", amount1)
	fmt.Printf("DAY 4 (2/2): %d\n", amount2)

	return nil
}
