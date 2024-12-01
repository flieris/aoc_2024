package main

import (
	"bufio"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type PageOrder struct {
	Page  int
	After []int
}

func getInputs(filePath string) ([]PageOrder, [][]int, error) {
	var pagesOrdering []PageOrder
	var manualPages [][]int
	fd, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			page1, _ := strconv.Atoi(parts[0])
			page2, _ := strconv.Atoi(parts[1])
			pageOrder := PageOrder{Page: page1, After: []int{page2}}
			index := slices.IndexFunc(pagesOrdering, func(p PageOrder) bool {
				return p.Page == pageOrder.Page
			})
			if index != -1 {
				pagesOrdering[index].After = append(pagesOrdering[index].After, pageOrder.After[0])
			} else {
				pagesOrdering = append(pagesOrdering, pageOrder)
			}
		} else if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			ints, err := StringsToIntegers(parts)
			if err != nil {
				return nil, nil, err
			}
			manualPages = append(manualPages, ints)
		}
	}
	err = scanner.Err()
	if err != nil {
		return nil, nil, err
	}
	return pagesOrdering, manualPages, nil
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

func getSumOfMiddlePages(pages [][]int) int {
	sum := 0
	for _, page := range pages {
		if len(page)%2 == 0 {
			log.Printf("Page: %v is even skipping\n", page)
			continue
		}
		sum += page[len(page)/2]
	}
	return sum
}

func getValidAndInvalidPages(pagesOrdering []PageOrder, manualPages [][]int) ([][]int, [][]int) {
	var validPages [][]int
	var invalidPages [][]int
	var correctPage int
	// Implement part 1 here
	for _, page := range manualPages {
		correctPage = 0
		for _, order := range pagesOrdering {
			tmp := 0
			// i could probaly use slices.Contains/Func here
			// but i want to also get an index of the element
			// so i can get the elements after and before it
			index := slices.IndexFunc(page, func(p int) bool {
				return p == order.Page
			})
			if index == -1 {
				continue
			}
			after := page[index+1:]
			before := page[:index]
			for _, afterPage := range order.After {
				if slices.Contains(before, afterPage) {
					tmp = -1
					break
				}
				if slices.Contains(after, afterPage) {
					tmp += 1
				}
			}
			if tmp > 0 {
				correctPage += 1
			}
			if tmp == -1 {
				// tainted page
				correctPage = -1
				break
			}
		}
		if correctPage > 0 {
			validPages = append(validPages, page)
		}
		if correctPage == -1 {
			invalidPages = append(invalidPages, page)
		}
	}
	return validPages, invalidPages
}

func part1(pagesOrdering []PageOrder, manualPages [][]int) int {
	var validPages [][]int
	validPages, _ = getValidAndInvalidPages(pagesOrdering, manualPages)
	return getSumOfMiddlePages(validPages)
}

func part2(pagesOrdering []PageOrder, manualPages [][]int) int {
	var invalidPages [][]int
	_, invalidPages = getValidAndInvalidPages(pagesOrdering, manualPages)
	for _, page := range invalidPages {
		// anonymous functions my beloved <3
		sort.Slice(page, func(i, j int) bool {
			left := page[i]
			right := page[j]
			index := slices.IndexFunc(pagesOrdering, func(p PageOrder) bool {
				return p.Page == left
			})
			if index == -1 {
				return false
			}
			order := pagesOrdering[index]
			return slices.Contains(order.After, right)
		})
	}
	return getSumOfMiddlePages(invalidPages)
}

func main() {
	pagesOrdering, manualPages, err := getInputs("day5/inputs.txt")
	if err != nil {
		panic(err)
	}
	solution1 := part1(pagesOrdering, manualPages)
	log.Printf("Solution for part 1: %d\n", solution1)
	solution2 := part2(pagesOrdering, manualPages)
	log.Printf("Solution for part 2: %d\n", solution2)
}
