package challenges

import (
	year2023 "github.com/lukebrobbs/advent-of-code/challenges/2023"
	year2024 "github.com/lukebrobbs/advent-of-code/challenges/2024"
)

type challengeMap map[int]func(input string, part2 bool) int

var YearChallenges = map[int]challengeMap{
	2023: year2023.Challenges,
	2024: year2024.Challenges,
}
