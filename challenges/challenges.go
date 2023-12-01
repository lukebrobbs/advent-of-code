package challenges

import (
	day1 "github.com/lukebrobbs/advent-of-code-2023/challenges/day-1"
)

type challenge func(input string, part2 bool) int

var DayChallenges = map[int]challenge{
	1: day1.Day1,
}
