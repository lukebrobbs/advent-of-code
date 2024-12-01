package day2

import (
	"strconv"
	"strings"
)

func getCubes() map[string]int {
	cubes := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return cubes
}

func Day2(input string, part2 bool) int {
	values := strings.Split(input, "\n")
	cubes := getCubes()
	validGames := []int{}
	power := []int{}

	for _, game := range values {
		gr := strings.Split(game, ":")
		gameNumber, _ := strconv.Atoi(strings.Split(gr[0], " ")[1])
		games := strings.Split(gr[1], ";")
		highestColors := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		valid := true
		for _, round := range games {
			picks := strings.Split(round, ",")
			for _, pick := range picks {
				pick = strings.TrimSpace(pick)
				total := strings.Split(pick, " ")
				amount, _ := strconv.Atoi(total[0])
				color := total[1]
				if highestColors[color] < amount {
					highestColors[color] = amount
				}
				if cubes[color] < amount {
					valid = false
				}
			}
		}
		if valid {
			validGames = append(validGames, gameNumber)
		}
		multiplied := highestColors["red"] * highestColors["green"] * highestColors["blue"]

		power = append(power, multiplied)
	}
	var total int

	if part2 {
		for _, v := range power {
			total += v
		}
	} else {
		for _, v := range validGames {
			total += v
		}
	}

	return total
}
