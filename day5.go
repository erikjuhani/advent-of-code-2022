package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Stacks map[int][]byte

type StacksMoveFn func(Stacks) func(int, int, int)

func MoveOne(s Stacks) func(int, int, int) {
	return func(amount, from, to int) {
		crates := reverse(s[from][len(s[from])-amount:])
		s[from] = s[from][:len(s[from])-amount]
		s[to] = append(s[to], crates...)
	}
}

func MoveAll(s Stacks) func(int, int, int) {
	return func(amount, from, to int) {
		crates := s[from][len(s[from])-amount:]
		s[from] = s[from][:len(s[from])-amount]
		s[to] = append(s[to], crates...)
	}
}

func reverse(b []byte) []byte {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return b
}

func ParseCommand(rawCmd string) (amount, from, to int) {
	cmds := []int{}
	for _, p := range strings.Split(rawCmd, " ") {
		d, err := strconv.Atoi(p)
		if err != nil {
			continue
		}
		cmds = append(cmds, d)
	}

	return cmds[0], cmds[1] - 1, cmds[2] - 1
}

func StacksFromInput(input []string) Stacks {
	stacks := make(Stacks)

	for _, r := range input {
		var i int
		for j := 0; j < len(r); j += 2 {
			if i%9 == 0 {
				i = 0
			}

			b := r[j]

			if b != 32 {
				stacks[i] = append([]byte{b}, stacks[i]...)
			}

			i++
		}
	}

	return stacks
}

func Solve(stacks Stacks, commands []string, fn StacksMoveFn) string {
	for _, c := range commands {
		fn(stacks)(ParseCommand(c))
	}

	keys := make([]int, 0)

	for k := range stacks {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var answer string
	for _, k := range keys {
		answer += string(reverse(stacks[k])[0])
	}

	return answer
}

func Day5() error {
	input, err := ReadInput("./input/day05")
	if err != nil {
		return err
	}

	schema := strings.Split(input, "\n\n")
	bytes := []byte{}

	for i := 1; i < len(schema[0]); i += 2 {
		bytes = append(bytes, schema[0][i])
	}

	d, c := strings.Split(string(bytes), "\n"), strings.Split(schema[1], "\n")

	data := d[:len(d)-1]
	commands := c[:len(c)-1]

	fmt.Printf("DAY 5 (1/2): %s\n", Solve(StacksFromInput(data), commands, MoveOne))
	fmt.Printf("DAY 5 (2/2): %s\n", Solve(StacksFromInput(data), commands, MoveAll))

	return nil
}
