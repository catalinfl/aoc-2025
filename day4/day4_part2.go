package day4

import "fmt"

func (s *Solution) Day4part2() {
	totalCount := 0
	for {
		counter := s.Day4part1()
		if counter == 0 {
			break
		}
		totalCount += counter
	}

	fmt.Println(totalCount)
}
