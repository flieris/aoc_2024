package main

import (
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getInputs(filePath string) ([]byte, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	output, err := io.ReadAll(fd)
	if err != nil {
		return nil, err
	}
	return output, err
}

func part1(inputs []byte) (int, error) {
	var output int

	mulReg, _ := regexp.Compile(`mul\(([0-9]{1,3},[0-9]{1,3})\)`)
	muls := mulReg.FindAllStringSubmatch(string(inputs), -1)
	for _, mul := range muls {
		vals := strings.Split(mul[1], ",")
		val1, _ := strconv.Atoi(vals[0])
		val2, _ := strconv.Atoi(vals[1])
		output += val1 * val2
	}

	return output, nil
}

func part2(inputs []byte) (int, error) {
	var output int

	mulReg, _ := regexp.Compile(`(do\(\)|don't\(\)|mul\(([0-9]{1,3},[0-9]{1,3})\))`)
	muls := mulReg.FindAllStringSubmatch(string(inputs), -1)
	log.Println(muls)
	run := true
	for _, mul := range muls {
		log.Println(mul)
		if strings.Contains(mul[0], "don't()") {
			run = false
			continue
		} else if strings.Contains(mul[0], "do()") {
			run = true
			continue
		}
		if !run {
			continue
		}
		vals := strings.Split(mul[2], ",")
		val1, _ := strconv.Atoi(vals[0])
		val2, _ := strconv.Atoi(vals[1])
		output += val1 * val2
	}

	return output, nil
}

func main() {
	inputs, err := getInputs("day3/inputs.txt")
	if err != nil {
		log.Fatalf("Error reading inputs: %v", err)
	}
	part1, err := part1(inputs)
	if err != nil {
		log.Fatalf("Error while doing part1: %v", err)
	}
	log.Printf("Solution for part 1: %d", part1)
	part2, err := part2(inputs)
	if err != nil {
		log.Fatalf("Error while doing part2: %v", err)
	}
	log.Printf("Solution for part2: %d", part2)
}
