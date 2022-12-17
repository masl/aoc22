package main

import (
	"fmt"
	"github.com/masl/aoc22/utils"
	"unicode"
)

func main() {
	input := utils.ReadInput(3)
	prioritySum := 0
	for _, line := range input {
		left := []rune(line)[:len(line)/2]
		right := []rune(line)[len(line)/2:]
		m := make(map[rune]bool)

		for _, char := range right {
			m[char] = findInSlice(left, char)
		}

		if k, ok := getTrueKey(m); ok {
			prioritySum += calculatePriority(k)
		}
	}

	fmt.Println("Part 1:", prioritySum)

	// Part 2
	prioritySum = 0
	for i := 0; i < len(input); i += 3 {
		first := []rune(input[i])
		second := []rune(input[i+1])
		third := []rune(input[i+2])

		m := make(map[rune]bool)
		for _, char := range first {
			m[char] = findInSlice(second, char) && findInSlice(third, char)
		}

		if k, ok := getTrueKey(m); ok {
			prioritySum += calculatePriority(k)
		}
	}

	fmt.Println("Part 2:", prioritySum)
}

func findInSlice(slice []rune, key rune) bool {
	for _, v := range slice {
		if key == v {
			return true
		}
	}

	return false
}

func getTrueKey(m map[rune]bool) (key rune, ok bool) {
	for k, v := range m {
		if v == true {
			key = k
			ok = true
			return
		}
	}

	return
}

func calculatePriority(char rune) (priority int) {
	if unicode.IsLower(char) {
		priority = int(char) - 96
		return
	}

	priority = int(char) - 38
	return
}
