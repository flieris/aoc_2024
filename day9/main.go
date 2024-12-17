package main

import (
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Block struct {
	Id         string
	BlockStart int
	Length     int
}

func getInputs(path string) []string {
	var output []string
	fd, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer fd.Close()
	content, err := io.ReadAll(fd)
	if err != nil {
		return nil
	}
	discString := strings.Split(strings.TrimSpace(string(content)), "")

	fileId := 0
	for id, val := range discString {
		intVal, _ := strconv.Atoi(val)
		if id%2 == 1 {
			if intVal == 0 {
				continue
			}
			for i := 0; i < intVal; i++ {
				output = append(output, ".")
			}
		} else {
			for i := 0; i < intVal; i++ {
				output = append(output, strconv.Itoa(fileId))
			}
			fileId++
		}

	}
	return output
}

func getLastBlockIndex(in []string, val string) int {
	last := 0
	for i := len(in) - 1; i >= 0; i-- {
		if in[i] == val {
			last = i
			break
		}
	}
	return last
}

func part1(inputs []string) int {
	checksum := 0
	defragedDisc := slices.Clone(inputs)
	for i := len(inputs) - 1; i >= 0; i-- {
		if inputs[i] == "." {
			continue
		}
		for j, block := range defragedDisc {
			if block == "." {
				defragedDisc = slices.Replace(defragedDisc, j, j+1, inputs[i])
				indexOfBlock := getLastBlockIndex(defragedDisc, inputs[i])
				defragedDisc = slices.Delete(defragedDisc, indexOfBlock, indexOfBlock+1)
				break
			}
		}
	}

	for i, block := range defragedDisc {
		blockValue, _ := strconv.Atoi(block)
		checksum += i * blockValue
	}
	return checksum
}

func mapBlocks(inputs []string) ([]Block, []Block) {
	files := make([]Block, 0)
	freeSpace := make([]Block, 0)
	for i, val := range inputs {
		if val == "." {
			if i == 0 {
				freeSpace = append(freeSpace, Block{val, i, 1})
				continue
			}
			if inputs[i-1] == "." {
				freeSpace[len(freeSpace)-1].Length++
				continue
			}
			freeSpace = append(freeSpace, Block{val, i, 1})
			continue
		} else {
			if i == 0 {
				files = append(files, Block{val, i, 1})
				continue
			}
			if val == inputs[i-1] {
				files[len(files)-1].Length++
				continue
			}
			files = append(files, Block{val, i, 1})
		}
	}
	return files, freeSpace
}

func part2(inputs []string) int {
	checksum := 0
	fileBlocks, freeBlocks := mapBlocks(inputs)
	for i := len(fileBlocks) - 1; i >= 0; i-- {
		for j, block := range freeBlocks {
			// don't move files to blocks that are after the file block
			if block.BlockStart > fileBlocks[i].BlockStart {
				break
			}
			if fileBlocks[i].Length <= block.Length {
				fileBlocks[i].BlockStart = block.BlockStart
				freeBlocks[j].BlockStart += fileBlocks[i].Length
				freeBlocks[j].Length -= fileBlocks[i].Length
				if freeBlocks[j].Length == 0 {
					freeBlocks = slices.Delete(freeBlocks, j, j+1)
				}
				break
			}
		}
	}
	for _, block := range fileBlocks {
		blockValue, _ := strconv.Atoi(block.Id)
		checksum += blockValue * (block.BlockStart*block.Length + (block.Length * (block.Length - 1) / 2))

	}
	return checksum
}

func main() {
	inputs := getInputs("day9/inputs.txt")
	solution1 := part1(inputs)
	log.Printf("Solution 1: %v", solution1)
	solution2 := part2(inputs)
	log.Printf("Solution 2: %v", solution2)
}
