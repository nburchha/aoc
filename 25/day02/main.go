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
	isRepeating := false
	for i := 1; i <= len(id)/2; i++ {
		if len(id)%i != 0 {
			continue
		}
		isRepeating1 := true
		for j := 1; j*i <= len(id); j++ {
			// fmt.Println("comparing: ", id[:i], id[i*(j-1):i*j], id)
			if id[:i] != id[i*(j-1):i*j] {
				isRepeating1 = false
				break
			}
		}
		if isRepeating1 {
			isRepeating = true
			break
		}
	}
	if !isRepeating {
		// fmt.Println("valid id: ", id)
		return 0
	}
	res, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// fmt.Println("invalid id: ", id)
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
