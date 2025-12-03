package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() []int {
	file, err := os.Open("../input/day01.txt")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var result []int
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if line[0] == 'L' {
			value = -value
		}
		result = append(result, value)
	}
	return result
}

func calcZeroPos(input []int) (int, int) {
	currentPos := 50
	res1, res2 := 0, 0
	for _, value := range input {
		// fmt.Println("currentPos: ", currentPos, "value: ", value)
		currentPos += value
		if currentPos <= 0 && currentPos-value != 0 {
			// fmt.Println("we add one to res2 because currentPos: ", currentPos, "at index: ", count)
			res2++
		}
		res2 += abs(currentPos / 100)
		fmt.Println(currentPos, abs(currentPos/100))
		currentPos = currentPos % 100
		if currentPos < 0 {
			currentPos += 100
		}
		if currentPos == 0 {
			res1++
		}
	}
	return res1, res2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input := parseInput()
	// fmt.Println(input)
	res1, res2 := calcZeroPos(input)
	fmt.Println("first exercise: ", res1)
	fmt.Println("second exercise: ", res2)
}
