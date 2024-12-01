package day1

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func countOccurrences(needle int, haystack []int) int {
	count := 0
	for _, num := range haystack {
		if num == needle {
			count++
		}
	}
	return count
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
	total := 0

	if part2 {
		for _, num := range list1 {
			total += num * countOccurrences(num, list2)
		}
		return total
	}

	sort.Ints(list1)
	sort.Ints(list2)
	for i, num := range list1 {
		total += int(math.Abs(float64(num - list2[i])))
	}
	return total
}
