package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getData(filePath string) ([][]int, error) {
	var output [][]int
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		ints, err := StringsToIntegers(parts)
		if err != nil {
			return nil, err
		}
		output = append(output, ints)
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func StringsToIntegers(input []string) ([]int, error) {
	output := make([]int, 0, len(input))
	for _, val := range input {
		n, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}
		output = append(output, n)
	}
	return output, nil
}

func sortingOrder(a, b int) int {
	tmp := a - b
	if tmp >= -3 && tmp <= 3 {
		return 1
	}
	return -1
}

func checkSlice(slice []int) bool {
	ascOrder := slices.IsSortedFunc(slice, func(a, b int) int {
		if a < b {
			return 1
		} else {
			return -1
		}
	})
	dscOrder := slices.IsSortedFunc(slice, func(a, b int) int {
		if a > b {
			return 1
		} else {
			return -1
		}
	})
	sorted := slices.IsSortedFunc(slice, sortingOrder)
	return (ascOrder || dscOrder) && sorted
}

func part1(reports [][]int) int {
	output := 0
	for _, report := range reports {
		if checkSlice(report) {
			output++
		}
	}
	return output
}
func part2(reports [][]int) int {
	output := 0
	for _, report := range reports {
		if checkSlice(report) {
			output++
			continue
		}
		for i := 0; i < len(report); i++ {
			test := slices.Clone(report)
			if i == len(test)-1 {
				test = test[:i]
			} else {
				test = append(test[:i], test[i+1:]...)
			}
			if checkSlice(test) {
				output++
				break
			}
		}
	}
	return output
}
func main() {
	reports, err := getData("day2/inputs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	output := part1(reports)
	log.Printf("Solution for part 1: %d\n", output)
	output2 := part2(reports)
	log.Printf("Soluttion for part 2: %d\n", output2)
}
