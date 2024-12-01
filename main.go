package main

import (
	"flag"
	"fmt"

	"github.com/lukebrobbs/advent-of-code/challenges"
	"github.com/lukebrobbs/advent-of-code/utils"
)

var year = flag.Int("year", 2023, "Advent of code year to run")
var day = flag.Int("day", 1, "Advent of code day to run")
var part2 = flag.Bool("part-2", false, "Should we run part 2 of the code")

func main() {
	flag.Parse()
	challenges := challenges.YearChallenges[*year]
	bytes, err := utils.FileReader(*year, *day)
	utils.Check(err)
	fmt.Println(challenges[*day](string(bytes), *part2))
}
