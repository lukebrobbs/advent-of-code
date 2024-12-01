package day1

import (
	"sort"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Day1(input string, part2 bool) (o int) {

	var list1 []int
	var list2 []int

	items := strings.Split(input, "\n")
	for _, item := range items {
		parts := strings.Split(item, "   ")
		num1, _ := strconv.Atoi(parts[0])
		num2, _ := strconv.Atoi(parts[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	// Sort both lists
	sort.Ints(list1)
	sort.Ints(list2)

	// Calculate differences between corresponding positions
	total := 0
	for i := 0; i < len(list1); i++ {
		total += abs(list1[i] - list2[i])
	}

	return total
}
