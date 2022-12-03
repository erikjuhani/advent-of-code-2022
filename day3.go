package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Compartment map[rune]struct{}

func (c Compartment) Has(r rune) bool {
	if _, ok := c[r]; ok {
		return true
	}
	return false
}

func RunePriority(r rune) int {
	mod := 0
	if unicode.IsLower(r) {
		mod = 96
	} else {
		mod = 38
	}
	return int(r) - mod
}

func compartment(input []rune) Compartment {
	var m = make(map[rune]struct{})
	for _, r := range input {
		m[r] = struct{}{}
	}
	return m
}

func Part1(input string) int {
	var sum int
	for _, rs := range strings.Split(input, "\n") {
		h := len(rs) / 2
		c0, c1 := compartment([]rune(rs[:h])), compartment([]rune(rs[h:]))

		for r := range c0 {
			if _, ok := c1[r]; ok {
				sum += RunePriority(r)
			}
		}
	}
	return sum
}

func Part2(input string) int {
	var (
		sum   int
		group []Compartment
	)

	for _, rs := range strings.Split(input, "\n") {
		if len(group) == 3 {
			for r := range group[0] {
				if group[1].Has(r) && group[2].Has(r) {
					sum += RunePriority(r)
				}
			}
			group = []Compartment{}
		}
		group = append(group, compartment([]rune(rs)))
	}

	return sum
}

func Day3() error {
	input, err := ReadInput("./input/day03")
	if err != nil {
		return err
	}

	fmt.Printf("DAY 3 (1/2): %d\n", Part1(input))
	fmt.Printf("DAY 3 (2/2): %d\n\n", Part2(input))

	return nil
}
