package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"slices"
)

type Point struct {
	x, y int
	val  rune
}

func getInputs(filePath string) ([][]Point, error) {
	var output [][]Point
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	reader := bufio.NewReader(fd)
	var line []Point
	lineCount := 0
	index := 0
	for {
		char, _, err := reader.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if char == '\n' {
			output = append(output, line)
			line = nil
			index = 0
			lineCount++
			continue
		}
		line = append(line, Point{x: index, y: lineCount, val: char})
		index++
	}

	return output, err
}

func getColumn(inputs [][]Point, index int) []Point {
	var column []Point
	for _, line := range inputs {
		column = append(column, line[index])
	}
	return column
}

func getColumns(inputs [][]Point, startIndex int, endIndex int) []Point {
	var column []Point
	for _, line := range inputs {
		column = append(column, line[startIndex:endIndex]...)
	}
	return column
}

func findSchema(line []Point, schema []Point) int {
	tmp := 0
	for _, char := range schema {
		if slices.Contains(line, char) {
			tmp += 1
		}
	}
	if tmp == len(schema) {
		return 1
	}
	return -1
}

func part1(inputs [][]Point) (int, error) {
	var count int
	directions := []Point{
		{x: 1, y: 0},
		{x: -1, y: 0},
		{x: 0, y: -1},
		{x: 0, y: 1},
		{x: 1, y: -1},
		{x: -1, y: -1},
		{x: 1, y: 1},
		{x: -1, y: 1},
	}
	for i, line := range inputs {
		for j, char := range line {
			if char.val == 'X' {
				for _, direction := range directions {
					schema := []Point{
						{x: j + direction.x, y: i + direction.y, val: 'M'},
						{x: j + 2*direction.x, y: i + 2*direction.y, val: 'A'},
						{x: j + 3*direction.x, y: i + 3*direction.y, val: 'S'},
					}
					if j+3*direction.x >= 0 && j+3*direction.x < len(line) && i+3*direction.y >= 0 && i+3*direction.y < len(inputs) {
						if direction.y == 0 {
							result := findSchema(line, schema)
							if result == 1 {
								count++
							}
						} else if direction.x == 0 {
							result := findSchema(getColumn(inputs, j), schema)
							if result == 1 {
								count++
							}
						} else {
							result := 0
							if direction.x < 0 {
								result = findSchema(getColumns(inputs, j+3*direction.x, j+1*direction.x+1), schema)
							} else {
								result = findSchema(getColumns(inputs, j+1*direction.x, j+3*direction.x+1), schema)
							}
							if result == 1 {
								count++
							}
						}
					}
				}
			}
		}
	}
	return count, nil
}

func part2(inputs [][]Point) (int, error) {
	var count int
	directions := [][]Point{
		{{x: -1, y: -1}, {x: 1, y: -1}},
		{{x: 1, y: -1}, {x: 1, y: 1}},
		{{x: -1, y: 1}, {x: 1, y: 1}},
		{{x: -1, y: -1}, {x: -1, y: 1}},
	}
	for i, line := range inputs {
		if i == 0 {
			continue
		}
		for j, char := range line {

			if char.val == 'A' {
				for _, direction := range directions {
					schema := []Point{
						{x: j + direction[0].x, y: i + direction[0].y, val: 'M'},
						{x: j + direction[1].x, y: i + direction[1].y, val: 'M'},
						{x: j - direction[0].x, y: i - direction[0].y, val: 'S'},
						{x: j - direction[1].x, y: i - direction[1].y, val: 'S'},
					}
					if j-1 >= 0 && j+2 <= len(line) {
						result := findSchema(getColumns(inputs, j-1, j+2), schema)
						if result == 1 {
							count++
							break
						}
					}
				}
			}
		}
	}

	return count, nil
}

func main() {
	test, err := getInputs("day4/inputs.txt")
	if err != nil {
		log.Fatalf("Error reading inputs: %v", err)
	}
	solution1, err := part1(test)
	if err != nil {
		log.Fatalf("Error while doing part1: %v", err)
	}
	log.Printf("Solution 1: %v", solution1)
	solution2, err := part2(test)
	if err != nil {
		log.Fatalf("Error while doing part2: %v", err)
	}
	log.Printf("Solution 2: %v", solution2)
}
