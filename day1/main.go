package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getInputs() ([]int, []int, error) {
	fd, err := os.Open("day1/inputs.txt")
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()
	var slice1, slice2 []int
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "   ")
		val1, err1 := strconv.Atoi(parts[0])
		val2, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			log.Println(err1, err2)
			continue
		}
		slice1 = append(slice1, val1)
		slice2 = append(slice2, val2)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	sort.Ints(slice1)
	sort.Ints(slice2)
	return slice1, slice2, nil
}

func GetTotalDistance(slice1, slice2 []int) int {
	var total_distance int
	for i := 0; i < len(slice1); i++ {
		distance := Abs(slice1[i] - slice2[i])
		total_distance += distance
	}
	return total_distance
}

func GetSimilalirytScore(slice1, slice2 []int) int {
	var similarity_score int
	for _, val1 := range slice1 {
		occurance := 0
		for _, val2 := range slice2 {
			if val1 == val2 {
				occurance += 1
			}
		}
		similarity_score += val1 * occurance
	}
	return similarity_score
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	slice1, slice2, err := getInputs()
	if err != nil {
		log.Fatal(err)
	}
	total_distance := GetTotalDistance(slice1, slice2)
	similarity_score := GetSimilalirytScore(slice1, slice2)
	log.Printf("Total distance: %d", total_distance)
	log.Printf("Similarity score: %d", similarity_score)
}
