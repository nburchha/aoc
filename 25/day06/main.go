package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.Join(strings.Fields(scanner.Text()), " "))
	}
	return lines
}

func main() {
	// lines := parseInput("../input/test06.txt")
	lines := parseInput("../input/day06.txt")
	fmt.Println(lines)
	symbols := strings.Split(lines[4], " ")
	fmt.Println(symbols)
	overallResult := 0
	for i, symbol := range symbols {
		res := 0
		for j := 0; j < len(lines)-1; j++ { // loop through the numbers
			split := strings.Split(lines[j], " ")
			tmp, _ := strconv.Atoi(split[i])
			fmt.Print(tmp, " ", string(symbol), " ")
			if j == 0 {
				res = tmp
			} else if symbol == "+" {
				res += tmp
			} else if symbol == "*" {
				res *= tmp
			}
		}
		fmt.Println("result:", res, "from index:", i)
		overallResult += res
	}
	fmt.Println(overallResult)
}
