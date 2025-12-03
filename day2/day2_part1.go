package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Day2part1() {
	distanceMap := getInput()
	var sum int64

	fmt.Println("Debug: starting Day2part1")
	for key, value := range distanceMap {
		fmt.Printf("Debug: range %d - %d\n", key, value)
		for i := key; i <= value; i++ {
			if verifyNumber(i) {
				fmt.Printf("Debug: found invalid ID %d\n", i)
				sum += i
			}
		}
	}

	fmt.Printf("Result: %d\n", sum)
}

func verifyNumber(num int64) bool {
	numStr := parseInt64ToString(num)
	n := len(numStr)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return numStr[:half] == numStr[half:]
}

func getInput() map[int64]int64 {
	file, err := os.ReadFile("input/day2.txt")
	if err != nil {
		panic(err)
	}

	txt := string(file)
	result := make(map[int64]int64)
	parts := strings.Split(txt, ",")
	for _, dist := range parts {
		dist = strings.TrimSpace(dist)
		if dist == "" {
			continue
		}
		split := strings.SplitN(dist, "-", 2)
		firstNum := parseStringToInt64(split[0])
		secondNum := parseStringToInt64(split[1])
		result[firstNum] = secondNum
	}

	return result
}

func parseInt64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func parseStringToInt64(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return num
}
