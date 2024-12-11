package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Stone struct {
	Val   int
	Times int
}

func getInputs(file string) ([]int, error) {
	output := make([]int, 0)

	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		for _, val := range parts {
			intval, _ := strconv.Atoi(val)
			output = append(output, intval)
		}
	}
	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func part1(inputs []int) int {
	for range 25 {
		copyInputs := []int{}
		for i := range inputs {
			length := len(strconv.Itoa(inputs[i]))
			if inputs[i] == 0 {
				copyInputs = append(copyInputs, 1)
			} else if length%2 == 0 {
				numStr := strconv.Itoa(inputs[i])
				part1, _ := strconv.Atoi(numStr[:length/2])
				part2, _ := strconv.Atoi(numStr[length/2:])
				copyInputs = append(copyInputs, part1, part2)

			} else {
				copyInputs = append(copyInputs, inputs[i]*2024)
			}
		}
		inputs = copyInputs
	}
	return len(inputs)
}

func parse(number int, timer int, counter *map[Stone]int) int {
	output := 0
	if timer > 0 {
		if val, ok := (*counter)[Stone{Val: number, Times: timer}]; ok {
			output += val
		} else {
			length := len(strconv.Itoa(number))
			if number == 0 {
				output += parse(1, timer-1, counter)
			} else if length%2 == 0 {
				numStr := strconv.Itoa(number)
				part1, _ := strconv.Atoi(numStr[:length/2])
				part2, _ := strconv.Atoi(numStr[length/2:])
				output += parse(part1, timer-1, counter)
				output += parse(part2, timer-1, counter)
			} else {
				output += parse(number*2024, timer-1, counter)
			}
			(*counter)[Stone{Val: number, Times: timer}] = output
		}
	} else {
		return 1
	}
	return output
}

func part2(inputs []int) int {
	counter := &map[Stone]int{}
	output := 0
	for i := range inputs {
		output += parse(inputs[i], 75, counter)
	}
	return output
}

func main() {
	inputs, err := getInputs("day11/inputs.txt")
	if err != nil {
		panic(err)
	}
	log.Printf("Inputs: %v", inputs)
	solution1 := part1(inputs)
	log.Printf("Solution 1: %v", solution1)
	solution2 := part2(inputs)
	log.Printf("Solution 2: %v", solution2)
}
