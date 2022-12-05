package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/masl/aoc22/utils"
)

func main() {
	input := utils.ReadInput(1)
	elfCount := 0
	elves := make(map[int]int)

	// Sums lines into a map
	for _, line := range input {
		if line == "" {
			elfCount++
		} else {
			num, err := strconv.Atoi(line)
			utils.ErrorCheck(err)

			_, ok := elves[elfCount]
			if ok {
				elves[elfCount] += num
			} else {
				elves[elfCount] = num
			}
		}
	}

	// Sort sums
	var firstSum []int
	for _, value := range elves {
		firstSum = append(firstSum, value)
	}

	// Part 1
	sort.Ints(firstSum)
	fmt.Println("Part 1:", firstSum[len(firstSum)-1])

	// Part 2
	secondSum := 0
	for _, v := range firstSum[len(firstSum)-3:] {
		secondSum += v
	}
	fmt.Println("Part 2:", secondSum)
}
