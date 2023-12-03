package challenges

import (
	day1 "github.com/lukebrobbs/advent-of-code-2023/challenges/day-1"
	day2 "github.com/lukebrobbs/advent-of-code-2023/challenges/day-2"
	day3 "github.com/lukebrobbs/advent-of-code-2023/challenges/day-3"
)

type challenge func(input string, part2 bool) int

var DayChallenges = map[int]challenge{
	1: day1.Day1,
	2: day2.Day2,
	3: day3.Day3,
}
