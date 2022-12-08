package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/erikjuhani/advent-of-code-2022/utils"
)

type Node struct {
	Parent   *Node
	Children map[string]*Node
	size     int
}

func (n *Node) Find(name string) *Node {
	for nn, node := range n.Children {
		if nn == name {
			return node
		}
	}
	return nil
}

func findNodes(node *Node) []*Node {
	var nodes []*Node
	for _, n := range node.Children {
		if n.size >= 8381165 {
			nodes = append(nodes, n)
			nodes = append(nodes, findNodes(n)...)
		}
	}
	return nodes
}

func traverse(node *Node) int {
	var sum int
	for _, n := range node.Children {
		sum += traverse(n)
		if n.size > 0 && n.size <= 100000 {
			sum += n.size
		}
	}
	return sum
}

func AddSize(size int, node *Node) {
	currentNode := node
	for currentNode != nil {
		currentNode.size += size
		currentNode = currentNode.Parent
	}
}

func Day7() error {
	input, err := utils.ReadInput("./input/day07")
	if err != nil {
		return err
	}

	commands := strings.Split(input, "\n")

	var root = &Node{Children: make(map[string]*Node)}
	currentNode := root
	for _, cmd := range commands[:len(commands)-1] {
		c := strings.Split(cmd, " ")
		switch c[0] {
		case "$":
			if c[1] == "cd" {
				dirname := c[2]
				if dirname == ".." {
					currentNode = currentNode.Parent
					continue
				}
				if n := currentNode.Find(dirname); n != nil {
					currentNode = currentNode.Find(dirname)
				}
			}
		case "dir":
			dirname := c[1]
			node := &Node{Parent: currentNode, Children: make(map[string]*Node)}
			currentNode.Children[dirname] = node
		default:
			size, err := strconv.Atoi(c[0])
			if err != nil {
				return err
			}

			AddSize(size, currentNode)
		}
	}

	sizes := make([]int, 0)

	for _, n := range findNodes(root) {
		sizes = append(sizes, n.size)
	}

	sort.Ints(sizes)

	fmt.Printf("DAY 7 (1/2): %d\n", traverse(root))
	fmt.Printf("DAY 7 (2/2): %+v\n", sizes[0])

	return nil
}
