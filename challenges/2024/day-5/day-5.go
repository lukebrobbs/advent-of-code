package day5

import (
	"strconv"
	"strings"
)

type Rule struct {
	Before []int
	After  []int
}

func indexOf[T comparable](element T, data []T) int {
	for i, v := range data {
		if v == element {
			return i
		}
	}
	return -1
}

func Day5(input string, part2 bool) int {
	parts := strings.Split(input, "\n\n")
	ordering := strings.Split(strings.TrimSpace(parts[0]), "\n")
	updates := strings.Split(strings.TrimSpace(parts[1]), "\n")
	var output [][]string
	total := 0

	var rules = make(map[int]Rule)

	for _, order := range ordering {
		ba := strings.Split(order, "|")
		before, _ := strconv.Atoi(ba[0])
		after, _ := strconv.Atoi(ba[1])

		beforeRule := rules[before]
		beforeRule.Before = append(beforeRule.Before, after)
		rules[before] = beforeRule

		afterRule := rules[after]
		afterRule.After = append(afterRule.After, before)
		rules[after] = afterRule
	}

	// iterate over updates
	// check the rules map, and ensure evertthing in before is before it, and everything in after is after
updates:
	for _, u := range updates {
		updateRow := strings.Split(u, ",")
		for i, update := range updateRow {

			num, _ := strconv.Atoi(update)

			for _, z := range rules[num].Before {
				index := indexOf(strconv.Itoa(z), updateRow)
				isValid := index == -1 || index > i
				if !isValid {
					continue updates
				}
			}
			for _, z := range rules[num].After {
				index := indexOf(strconv.Itoa(z), updateRow)
				isValid := index == -1 || index < i
				if !isValid {
					continue updates
				}
			}
		}
		output = append(output, updateRow)
	}

	for _, n := range output {
		middleItem := n[len(n)/2]
		num, _ := strconv.Atoi(middleItem)
		total += num
	}

	return total
}
