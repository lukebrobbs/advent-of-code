package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type Handler = func(n1, n2 int) bool

const (
	GREATER_THAN = "gt"
	LESS_THAN    = "lt"
	MIN_DIFF     = 1
	MAX_DIFF     = 3
)

func isSafeDiff(diff int) bool {
	return diff >= MIN_DIFF && diff <= MAX_DIFF
}

var handler = map[string]Handler{
	GREATER_THAN: func(n1, n2 int) bool {
		return isSafeDiff(n2 - n1)
	},
	LESS_THAN: func(n1, n2 int) bool {
		return isSafeDiff(n1 - n2)
	},
}

func determineRelationship(numbers []int) (string, error) {
	if numbers[0] >= numbers[1] {
		return LESS_THAN, nil
	}
	if numbers[0] <= numbers[1] {
		return GREATER_THAN, nil
	}

	return "", fmt.Errorf("neither strictly increasing nor decreasing")
}

func isSequenceSafe(numbers []int) bool {
	rel, err := determineRelationship(numbers)
	if err != nil {
		return false
	}

	for i := 0; i < len(numbers)-1; i++ {
		if !handler[rel](numbers[i], numbers[i+1]) {
			return false
		}
	}
	return true
}

func Day2(input string, part2 bool) int {
	items := strings.Split(input, "\n")
	safeCount := 0

	for _, input := range items {
		fields := strings.Fields(input)
		numbers := make([]int, len(fields))
		for i, field := range fields {
			numbers[i], _ = strconv.Atoi(field)
		}

		// Check if sequence is safe without removing any numbers
		if isSequenceSafe(numbers) {
			safeCount++
			continue
		}

		// Part 2: Try removing each number one at a time
		if part2 {
			for i := range numbers {
				// Create a new slice without the current number
				tempNumbers := make([]int, 0, len(numbers)-1)
				tempNumbers = append(tempNumbers, numbers[:i]...)
				tempNumbers = append(tempNumbers, numbers[i+1:]...)

				if isSequenceSafe(tempNumbers) {
					safeCount++
					break
				}
			}
		}
	}
	return safeCount
}
