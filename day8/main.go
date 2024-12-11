package main

import (
	"bufio"
	"log"
	"os"
)

type Node struct {
	X, Y int
	Freq rune
}

func getInputs(filePath string) ([][]Node, error) {
	var output [][]Node
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineSlice := []Node{}
		for i, val := range line {
			lineSlice = append(lineSlice, Node{i, j, val})
		}
		output = append(output, lineSlice)
		j++
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return output, nil
}

func findNodes(inputs [][]Node) []Node {
	var output []Node
	for _, row := range inputs {
		for _, node := range row {
			if node.Freq != '.' {
				output = append(output, node)
			}
		}
	}
	return output
}

func GetAntinode(node1, node2 Node) Node {
	distX := node1.X - node2.X
	distY := node1.Y - node2.Y

	return Node{node1.X + distX, node1.Y + distY, '#'}
}

func GetAntinodes(node1, node2 Node, row int, col int) []Node {
	output := []Node{}
	distX := node1.X - node2.X
	distY := node1.Y - node2.Y
	next := Node{node1.X + distX, node1.Y + distY, '#'}
	for !OutOfBounds(next, row, col) {
		output = append(output, next)
		next.X += distX
		next.Y += distY
	}
	next = Node{node1.X - distX, node1.Y - distY, '#'}
	for !OutOfBounds(next, row, col) {
		output = append(output, next)
		next.X -= distX
		next.Y -= distY

	}
	return output
}

func GetSimilarFreq(node Node, nodes []Node) []Node {
	var output []Node
	for _, n := range nodes {
		if n.Freq == node.Freq && n != node {
			output = append(output, n)
		}
	}
	return output
}

func OutOfBounds(node Node, row, col int) bool {
	return node.X < 0 || node.X >= col || node.Y < 0 || node.Y >= row
}

func IsUnique(node Node, nodes []Node) bool {
	for _, n := range nodes {
		if n == node {
			return false
		}
	}
	return true
}

func part1(inputs [][]Node) int {
	antinodes := []Node{}
	col := len(inputs[0])
	row := len(inputs)
	nodes := findNodes(inputs)
	for node := range nodes {
		similarFreq := GetSimilarFreq(nodes[node], nodes)
		for _, n := range similarFreq {
			antinode := GetAntinode(nodes[node], n)
			if !OutOfBounds(antinode, row, col) && IsUnique(antinode, nodes) && IsUnique(antinode, antinodes) {
				antinodes = append(antinodes, antinode)
			}
		}
	}
	return len(antinodes)
}

func part2(inputs [][]Node) int {
	antinodes := []Node{}
	col := len(inputs[0])
	row := len(inputs)
	nodes := findNodes(inputs)
	for node := range nodes {
		similarFreq := GetSimilarFreq(nodes[node], nodes)
		for _, n := range similarFreq {
			resonants := GetAntinodes(nodes[node], n, row, col)
			for _, r := range resonants {
				if !OutOfBounds(r, row, col) && IsUnique(r, antinodes) {
					antinodes = append(antinodes, r)
				}
			}
		}

	}
	return len(antinodes)
}

func main() {
	inputs, err := getInputs("day8/inputs.txt")
	if err != nil {
		panic(err)
	}
	solution1 := part1(inputs)
	log.Printf("Solution 1: %v", solution1)
	solution2 := part2(inputs)
	log.Printf("Solution 2: %v", solution2)
}
