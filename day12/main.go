package main

import (
	"bufio"
	"log"
	"os"
	"time"
)

type Plants struct {
	plant rune
	x, y  int
}
type Region struct {
	plants []Plants
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

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

func inBounds(x, y int, inputs [][]rune) bool {
	if x < 0 || x >= len(inputs[0]) || y < 0 || y >= len(inputs) {
		return false
	}
	return true
}

func defineRegions(inputs [][]rune) []Region {
	regions := make([]Region, 0)
	visited := make([][]bool, len(inputs))
	for i := range visited {
		visited[i] = make([]bool, len(inputs[0]))
	}

	var floodFill func(x, y int, plant rune) []Plants
	floodFill = func(x, y int, plant rune) []Plants {
		if !inBounds(x, y, inputs) || visited[y][x] || inputs[y][x] != plant {
			return nil
		}
		visited[y][x] = true
		plants := []Plants{{plant: plant, x: x, y: y}}
		for _, dir := range directions {
			newX := x + dir[0]
			newY := y + dir[1]
			plants = append(plants, floodFill(newX, newY, plant)...)
		}
		return plants
	}

	for y := 0; y < len(inputs); y++ {
		for x := 0; x < len(inputs[0]); x++ {
			if !visited[y][x] {
				regionPlants := floodFill(x, y, inputs[y][x])
				if len(regionPlants) > 0 {
					regions = append(regions, Region{plants: regionPlants})
				}
			}
		}
	}

	return regions
}

func part1(inputs [][]rune) int {
	price := 0
	regions := defineRegions(inputs)
	for _, region := range regions {
		regionPrice := 0
		for _, plant := range region.plants {
			for _, dir := range directions {
				newX := plant.x + dir[0]
				newY := plant.y + dir[1]
				if inBounds(newX, newY, inputs) && inputs[newY][newX] != plant.plant {
					regionPrice++
				} else if !inBounds(newX, newY, inputs) {
					regionPrice++
				}
			}
		}
		price += len(region.plants) * regionPrice
	}

	return price
}

func part2(inputs [][]rune) int {
	price := 0
	regions := defineRegions(inputs)
	for _, region := range regions {
		regionPrice := 0
		for _, plant := range region.plants {
			for _, dir := range directions {
				newX := plant.x + dir[0]
				newY := plant.y + dir[1]
				if inBounds(newX, newY, inputs) && inputs[newY][newX] != plant.plant {
					regionPrice++
				} else if !inBounds(newX, newY, inputs) {
					regionPrice++
				}
			}
		}
		price += len(region.plants) * regionPrice
	}

	return price
}

func main() {
	inputs, _ := getInputs("day12/inputs.txt")
	start := time.Now()
	solution1 := part1(inputs)
	log.Printf("Time for part1: %v", time.Since(start))
	log.Printf("Solution 1: %v", solution1)
	start = time.Now()
	solution2 := part2(inputs)
	log.Printf("Time for part2: %v", time.Since(start))
	log.Printf("Solution 2: %v", solution2)
}
