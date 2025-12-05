package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput() ([][2]int, error) {
	file, err := os.Open("../input/day02.txt")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	splits := strings.Split(line, ",")
	var result [][2]int
	for _, split := range splits {
		nums := strings.Split(split, "-")
		if len(nums) == 2 {
			n1, _ := strconv.Atoi(nums[0])
			n2, _ := strconv.Atoi(nums[1])
			result = append(result, [2]int{n1, n2})
		}
	}
	return result, nil
}

func calculateInvalidIDs(id string) int {
	res := 0
	for i, _ := range id {
		// fmt.Println(i, id[:i], id[i:])
		if id[:i] == id[i:] {
			// fmt.Println("invalid id: ", id)
			tmp, err := strconv.Atoi(id)
			if err != nil {
				fmt.Println(err)
				return 0
			}
			res += tmp
		}
	}
	return res
}

func calculateInvalidIDs2(id string) int {
	if len(id) < 2 {
		return 0
	}
	doubled := id + id
	if !strings.Contains(doubled[1:len(doubled)-1], id) {
		return 0
	}
	res, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return res
}

func main() {
	input, err := parseInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(input)
	res1 := 0
	res2 := 0
	for _, idRange := range input {
		for i := idRange[0]; i <= idRange[1]; i++ {
			res1 += calculateInvalidIDs(strconv.Itoa(i))
			res2 += calculateInvalidIDs2(strconv.Itoa(i))
		}
	}

	fmt.Println(res1)
	fmt.Println(res2)
}
