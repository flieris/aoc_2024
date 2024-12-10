package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func evaluate(result int64, val int64, numbers []int64) bool {
	if len(numbers) == 0 {
		return val == result
	}
	if evaluate(result, val+numbers[0], numbers[1:]) {
		return true
	}
	return evaluate(result, val*numbers[0], numbers[1:])
}

func evaluate2(result int64, val int64, numbers []int64) bool {
	if len(numbers) == 0 {
		return val == result
	}
	if evaluate2(result, val+numbers[0], numbers[1:]) {
		return true
	}
	if evaluate2(result, val*numbers[0], numbers[1:]) {
		return true
	}
	tmp1 := strconv.FormatInt(val, 10)
	tmp2 := strconv.FormatInt(numbers[0], 10)
	concat := tmp1 + tmp2
	tmp, err := strconv.ParseInt(concat, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return evaluate2(result, tmp, numbers[1:])
}

func part1() {
	fd, err := os.Open("day7/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		result, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums := []int64{}
		for _, val := range strings.Fields(parts[1]) {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, num)
		}
		if evaluate(result, nums[0], nums[1:]) {
			sum += result
		}
	}
	log.Printf("Part 1: %d\n", sum)
}

func part2() {
	fd, err := os.Open("day7/inputs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, ":")
		result, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums := []int64{}
		for _, val := range strings.Fields(parts[1]) {
			num, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				log.Fatal(err)
			}

			nums = append(nums, num)
		}
		if evaluate2(result, nums[0], nums[1:]) {
			sum += result
		}
	}
	log.Printf("Part 2: %d\n", sum)
}

func main() {
	part1()
	part2()
}
