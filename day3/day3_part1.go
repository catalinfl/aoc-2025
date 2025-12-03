package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Day3part1() {
	banks := getInput()
	sum := 0
	for _, bank := range banks {
		bank = strings.TrimSpace(bank)
		biggestBank := -1
		indexBiggestBank := 0
		for i := 0; i < len(bank)-1; i++ {
			if biggestBank < int(bank[i]) {
				biggestBank = int(bank[i])
				indexBiggestBank = i
			}
		}

		secondBiggestBank := -1
		for i := indexBiggestBank + 1; i < len(bank); i++ {
			secondBiggestBank = max(secondBiggestBank, int(bank[i]))
		}

		digit1 := biggestBank - int('0')
		digit2 := secondBiggestBank - int('0')

		numStr := fmt.Sprintf("%d%d", digit1, digit2)
		fmt.Println(numStr)
		sum += atoi(numStr)
	}

	fmt.Println(sum)
}

func atoi(s string) int {
	num, _ := strconv.Atoi(strings.TrimSpace(s))
	return num
}

func getInput() []string {
	day3, err := os.ReadFile("input/day3.txt")
	if err != nil {
		panic(err)
	}

	txt := string(day3)
	return strings.Split(txt, "\n")
}
