package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

type Point struct {
	X, Y int
	Val  rune
}

func getInputs(filePath string) ([][]rune, error) {
	var output [][]rune
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	j := 0
	for scanner.Scan() {
		line := scanner.Text()
		output = append(output, []rune(line))
		j++
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}

	return output, nil
}

func getGuardStartingPos(inputs [][]rune) (int, int, int, int) {
	for i, line := range inputs {
		for j, val := range line {
			switch val {
			case '^':
				return j, i, 0, -1
			case 'v':
				return j, i, 0, 1
			case '<':
				return j, i, -1, 0
			case '>':
				return j, i, 1, 0
			}
		}
	}
	return 0, 0, 0, 0
}

func GuardRoutePrediction(inputs [][]rune, indexX int, indexY int, pos int, dirX int, dirY int) int {
	if indexX+dirX < 0 || indexX+dirX >= len(inputs[0]) || indexY+dirY < 0 || indexY+dirY >= len(inputs) {
		return 1
	}
	nextPos := inputs[indexY+dirY][indexX+dirX]
	moveX := indexX
	moveY := indexY
	if nextPos == '#' {
		tmp := dirX
		dirX = -1 * dirY
		dirY = tmp
	}
	inputs[indexY][indexX] = 'X'
	pos += GuardRoutePrediction(inputs, moveX+dirX, moveY+dirY, pos, dirX, dirY)

	return pos
}

func checkIfInBounds(inputs [][]rune, indexX int, indexY int) bool {
	if indexX < 0 || indexX >= len(inputs[0]) || indexY < 0 || indexY >= len(inputs) {
		return false
	}
	return true
}

func dirMapping(dirX int, dirY int) rune {
	if dirX == 0 && dirY == -1 {
		return '^'
	}
	if dirX == 0 && dirY == 1 {
		return 'v'
	}
	if dirX == -1 && dirY == 0 {
		return '<'
	}
	if dirX == 1 && dirY == 0 {
		return '>'
	}
	return ' '
}

func checkIfLooping(inputs [][]rune, indexX int, indexY int, dirX int, dirY int) bool {
	visited := make(map[Point]map[rune]bool)
	for checkIfInBounds(inputs, indexX+dirX, indexY+dirY) {
		key := Point{X: indexX, Y: indexY}
		dir := dirMapping(dirX, dirY)
		if _, ok := visited[key]; !ok {
			visited[key] = make(map[rune]bool)
		}
		if visited[key][dir] {
			return true
		}
		visited[key][dir] = true
		nextPos := inputs[indexY+dirY][indexX+dirX]
		if nextPos == '#' {
			tmp := dirX
			dirX = -1 * dirY
			dirY = tmp
		}
		indexX += dirX
		indexY += dirY

	}
	return false
}

func part1(inputs [][]rune) int {
	posX, posY, dirX, dirY := getGuardStartingPos(inputs)
	log.Printf("posX: %v, posY: %v, dirX: %v, dirY: %v", posX, posY, dirX, dirY)
	_ = GuardRoutePrediction(inputs, posX, posY, 1, dirX, dirY)
	distincPositions := 1
	for _, line := range inputs {
		for _, val := range line {
			if val == 'X' {
				distincPositions++
			}
		}
	}
	return distincPositions
}

func part2(inputs [][]rune) int {
	posX, posY, dirX, dirY := getGuardStartingPos(inputs)
	log.Printf("posX: %v, posY: %v, dirX: %v, dirY: %v", posX, posY, dirX, dirY)
	count := 0
	for checkIfInBounds(inputs, posX+dirX, posY+dirY) {
		nextPos := inputs[posY+dirY][posX+dirX]
		moveX := posX
		moveY := posY

		inputs[posY+dirY][posX+dirX] = '#'
		if checkIfLooping(inputs, moveX, moveY, dirX, dirY) {
			count++
		}
		inputs[posY+dirY][posX+dirX] = '.'
		if nextPos == '#' {
			tmp := dirX
			dirX = -1 * dirY
			dirY = tmp
		}
		posX = moveX + dirX
		posY = moveY + dirY

	}

	return count
}

func main() {
	inputs, err := getInputs("day6/inputs.txt")
	if err != nil {
		log.Fatalf("Error getting inputs: %v", err)
	}
	start := time.Now()
	solution1 := part1(inputs)
	log.Printf("solution for part 1: %v", solution1)
	log.Printf("Part 1 execution time: %v", time.Since(start))
	inputs, err = getInputs("day6/inputs.txt")
	if err != nil {
		log.Fatalf("Error getting inputs: %v", err)
	}
	start = time.Now()
	solution2 := part2(inputs)
	log.Printf("solution for part 2: %v", solution2)
	log.Printf("Part 2 execution time: %v", time.Since(start))
}
