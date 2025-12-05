package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	// "strings"
)

func parseInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func findLargestRune(s string) int {
	if len(s) == 0 {
		return 0
	}

	index := 0
	var maxRune rune = 0
	for i, r := range s {
		if r > maxRune {
			maxRune = r
			index = i
		}
	}

	return index
}

func getMaxBatteryCapacity(line string, batteryCount, res int) int {
	if batteryCount <= 0 {
		return res
	}
	index := findLargestRune(line[:len(line)-(batteryCount-1)])
	val, _ := strconv.Atoi(string(line[index]))
	return getMaxBatteryCapacity(line[index+1:], batteryCount-1, res*10+val)
}

func main() {
	input := parseInput("../input/day03.txt")
	res1 := 0
	res2 := 0
	for _, line := range input {
		fmt.Println(line)
		res1 += getMaxBatteryCapacity(line, 2, 0)
		tmp := getMaxBatteryCapacity(line, 12, 0)
		fmt.Println("result: ", tmp, "line: ", line)
		res2 += tmp
	}
	fmt.Println("1.: ", res1)
	fmt.Println("2.: ", res2)
}
