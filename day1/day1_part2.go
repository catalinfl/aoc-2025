package day1

import (
	"fmt"
	"os"
	"strings"
)

func (s *Solution) Day1part2() {
	g, err := os.ReadFile("input/day1.txt")
	if err != nil {
		panic(err)
	}

	replace := strings.ReplaceAll(string(g), "\r\n", "\n")
	str := strings.Split(replace, "\n")
	cypher := 50
	counter := 0

	for _, line := range str {
		num := convertStrToInt(line[1:])
		sign := line[0]
		for i := 0; i < num; i++ {
			cypher = (cypher + map[rune]int{'L': -1, 'R': 1}[rune(sign)] + 100) % 100
			if cypher == 0 {
				counter++
			}
		}
	}

	fmt.Println(counter)

}
