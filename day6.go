package main

import (
	"fmt"

	"github.com/erikjuhani/advent-of-code-2022/utils"
)

func subroutine(input string, length int) int {
	for i := range input {
		if len(set([]byte(input[i:i+length]))) == length {
			return i + length
		}
	}
	return 0
}

func set[T comparable](s []T) map[T]struct{} {
	m := make(map[T]struct{})
	for _, v := range s {
		m[v] = struct{}{}
	}
	return m
}

func Day6() error {
	input, err := utils.ReadInput("./input/day06")
	if err != nil {
		return err
	}

	fmt.Printf("DAY 6 (1/2): %d\n", subroutine(input, 4))
	fmt.Printf("DAY 6 (2/2): %d\n", subroutine(input, 14))

	return nil
}
