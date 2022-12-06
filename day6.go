package main

import "fmt"

func subroutine(input string, length int) int {
	for i := 0; i < len(input); i++ {
		j := i + length
		if j > len(input) {
			break
		}
		if charsDiff(input[i:j]) {
			return j
		}
	}

	return 0
}

func charsDiff(chars string) bool {
	m := make(map[rune]struct{})
	for _, r := range chars {
		if _, ok := m[r]; ok {
			return false
		}

		m[r] = struct{}{}
	}

	return true
}

func StartOfAPacket(input string) int {
	return subroutine(input, 4)
}

func StartOfAMessage(input string) int {
	return subroutine(input, 14)
}

func Day6() error {
	input, err := ReadInput("./input/day06")
	if err != nil {
		return err
	}

	fmt.Printf("DAY 6 (1/2): %d\n", StartOfAPacket(input))
	fmt.Printf("DAY 6 (2/2): %d\n", StartOfAMessage(input))

	return nil
}
