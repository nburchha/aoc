package main

import (
	"bufio"
	"fmt"
	"os"
	// "strconv"
	// "strings"
)

func parseInput(filename string) [][]bool {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var paperRolls [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		newRow := make([]bool, len(scanner.Text()))
		for x, rune := range scanner.Text() {
			if rune == '@' {
				newRow[x] = true
			}
		}
		paperRolls = append(paperRolls, newRow)
	}
	return paperRolls
}

func checkAdjacentPaperRolls(grid [][]bool, x, y int) bool {
	paperRollCount := 0
	for newY := y - 1; newY <= y+1; newY++ {
		if newY < 0 || newY >= len(grid) {
			continue
		}
		for newX := x - 1; newX <= x+1; newX++ {
			// fmt.Println(y, "|", x, ": ", newY, "|", newX)
			if newX < 0 || newX >= len(grid[newY]) {
				continue
			}
			if grid[newY][newX] && (newY != y || newX != x) {
				// fmt.Println(y, "|", x, ": ", newY, "|", newX, "<- this is a paperRoll")
				paperRollCount++
			}
		}
	}
	return paperRollCount < 4
}

// loops through grid and returns number of paperRolls that have < 4 adjacent paperRolls.
func loopGrid(grid [][]bool) [][2]int {
	var toRemove [][2]int
	for y, row := range grid {
		for x, val := range row {
			if val && checkAdjacentPaperRolls(grid, x, y) {
				// fmt.Println("< 4 paperRolls adjacent to them\n", y, x)
				toRemove = append(toRemove, [2]int{y, x})
			}
		}
	}
	return toRemove
}

func removePaperRolls(grid [][]bool, remove [][2]int) {
	for _, paperRoll := range remove {
		grid[paperRoll[0]][paperRoll[1]] = false
	}
}

func main() {
	input := parseInput("../input/day04.txt")
	fmt.Println("first puzzle: ", len(loopGrid(input)))
	// 2. puzzle
	res := 0
	grid := input
	for toRemove := loopGrid(grid); len(toRemove) > 0; toRemove = loopGrid(grid) {
		res += len(toRemove)
		removePaperRolls(grid, toRemove)
	}
	fmt.Println("second puzzle: ", res)
}
