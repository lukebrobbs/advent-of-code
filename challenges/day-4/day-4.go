package day4

import (
	"regexp"
	"slices"
	"strings"
)

func Day4(input string, part2 bool) (total int) {
	cards := strings.Split(input, "\n")
	for _, card := range cards {
		reg := regexp.MustCompile(`Card\s+\d+:`)
		res := reg.ReplaceAllString(card, "${1}")
		nums := strings.Split(res, "|")
		myNums := strings.Fields(nums[0])
		winningNums := strings.Fields(nums[1])
		totalWinningNums := 0
		t := 0

		for _, num := range myNums {
			if slices.Contains(winningNums, num) {
				totalWinningNums++
			}
		}

		for i := 0; i < totalWinningNums; i++ {
			if i == 0 {
				t += 1
			} else {
				t *= 2
			}
		}
		total += t
	}
	return
}
