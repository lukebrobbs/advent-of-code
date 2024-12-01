package challenges

import (
	year2023 "github.com/lukebrobbs/advent-of-code-2023/challenges/2023"
)

type challengeMap map[int]func(input string, part2 bool) int

var YearChallenges = map[int]challengeMap{
	2023: year2023.Challenges,
}
