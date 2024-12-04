package day4

import (
	"strings"
)

func Day4(input string, part2 bool) int {
	// Parse input into grid
	lines := strings.Split(strings.TrimSpace(input), "\n")
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	if !part2 {
		return findXMAS(grid)
	}
	return findXMASPattern(grid)
}

func findXMAS(grid [][]rune) int {
	// All possible directions
	directions := [][2]int{
		{0, 1}, {1, 1}, {1, 0}, {1, -1},
		{0, -1}, {-1, -1}, {-1, 0}, {-1, 1},
	}

	count := 0
	word := "XMAS"

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if checkWord(grid, i, j, dir, word) {
					count++
				}
			}
		}
	}
	return count
}

func findXMASPattern(grid [][]rune) int {
	count := 0

	// Check each position as potential center of X
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			// Check if this position can be center of X
			if checkXPattern(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func checkXPattern(grid [][]rune, centerRow, centerCol int) bool {
	// Diagonal directions for X pattern
	diagonals := [][2]int{
		{-1, -1}, // up-left
		{-1, 1},  // up-right
		{1, -1},  // down-left
		{1, 1},   // down-right
	}

	// Check each diagonal pair (forming the X)
	for i := 0; i < 4; i += 2 {
		dir1 := diagonals[i]
		dir2 := diagonals[i+1]

		// Get the three characters in each diagonal
		diag1 := []rune{
			grid[centerRow+dir1[0]][centerCol+dir1[1]],
			grid[centerRow][centerCol],
			grid[centerRow-dir1[0]][centerCol-dir1[1]],
		}

		diag2 := []rune{
			grid[centerRow+dir2[0]][centerCol+dir2[1]],
			grid[centerRow][centerCol],
			grid[centerRow-dir2[0]][centerCol-dir2[1]],
		}

		// Check if either diagonal contains "MAS" (forward or backward)
		if !isValidMAS(diag1) || !isValidMAS(diag2) {
			return false
		}
	}
	return true
}

func isValidMAS(chars []rune) bool {
	// Check forward
	if string(chars) == "MAS" {
		return true
	}
	// Check backward
	if string(chars) == "SAM" {
		return true
	}
	return false
}

func checkWord(grid [][]rune, row, col int, dir [2]int, word string) bool {
	if grid[row][col] != rune(word[0]) {
		return false
	}

	for i := 0; i < len(word); i++ {
		newRow := row + i*dir[0]
		newCol := col + i*dir[1]

		if newRow < 0 || newRow >= len(grid) ||
			newCol < 0 || newCol >= len(grid[0]) {
			return false
		}
	}

	for i := 0; i < len(word); i++ {
		if grid[row+i*dir[0]][col+i*dir[1]] != rune(word[i]) {
			return false
		}
	}

	return true
}
