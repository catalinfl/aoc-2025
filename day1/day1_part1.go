package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Day1part1() {
	g, err := os.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}

	replace := strings.ReplaceAll(string(g), "\r\n", "\n")
	gSplit := strings.Split(replace, "\n")

	cypher := 50
	counter := 0

	for _, line := range gSplit {
		sign := line[0]
		value := line[1:]
		switch sign {
		case 'L':
			cypher = (cypher - convertStrToInt(value) + 100) % 100
		case 'R':
			cypher = (cypher + convertStrToInt(value)) % 100
		}
		if cypher == 0 {
			counter++
		}

	}
	println(counter)
}

func convertStrToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
