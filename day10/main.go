package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var DIRECTIONS = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func parseInput(path string) [][]int {
	var output [][]int
	fd, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "")
		intLine := []int{}
		for _, part := range parts {
			val, _ := strconv.Atoi(part)
			intLine = append(intLine, val)
		}
		output = append(output, intLine)
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	return output
}

func dfs(inputs [][]int, x, y int, visited [][]bool, count *int) {
	if !isValid(x, y, inputs) {
		return
	}
	if visited[y][x] {
		return
	}
	visited[y][x] = true
	if inputs[y][x] == 9 {
		*count++
	}

	for _, dir := range DIRECTIONS {
		newX := x + dir[0]
		newY := y + dir[1]
		if isValid(newX, newY, inputs) && inputs[newY][newX] == inputs[y][x]+1 {
			dfs(inputs, newX, newY, visited, count)
		}
	}
}

func dfsRaiting(inputs [][]int, x, y int, count *int) {
	if !isValid(x, y, inputs) {
		return
	}
	if inputs[y][x] == 9 {
		*count++
	}

	for _, dir := range DIRECTIONS {
		newX := x + dir[0]
		newY := y + dir[1]
		if isValid(newX, newY, inputs) && inputs[newY][newX] == inputs[y][x]+1 {
			dfsRaiting(inputs, newX, newY, count)
		}
	}
}

func isValid(x, y int, inputs [][]int) bool {
	return x >= 0 && x < len(inputs[0]) && y >= 0 && y < len(inputs)
}

func part1(inputs [][]int) int {
	score := 0
	width := len(inputs[0])
	height := len(inputs)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if inputs[y][x] == 0 {
				visited := make([][]bool, height)
				for i := range visited {
					visited[i] = make([]bool, width)
				}
				counter := 0
				dfs(inputs, x, y, visited, &counter)
				score += counter
			}
		}
	}
	return score
}

func part2(inputs [][]int) int {
	score := 0
	width := len(inputs[0])
	height := len(inputs)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if inputs[y][x] == 0 {
				counter := 0
				dfsRaiting(inputs, x, y, &counter)
				score += counter
			}
		}
	}
	return score
}

func main() {
	inputs := parseInput("day10/inputs.txt")

	start := time.Now()
	solution1 := part1(inputs)
	log.Printf("Solution took %v", time.Since(start))
	log.Printf("Solution 1: %d", solution1)
	start = time.Now()
	solution2 := part2(inputs)
	log.Printf("Solution took %v", time.Since(start))
	log.Printf("Solution 2: %d", solution2)
}
