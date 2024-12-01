package year2023

import (
	day1 "github.com/lukebrobbs/advent-of-code/challenges/2024/day-1"
)

type challenge = func(input string, part2 bool) int

var Challenges = map[int]challenge{
	1: day1.Day1,
}
