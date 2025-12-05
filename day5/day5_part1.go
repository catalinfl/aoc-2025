package day5

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Solution struct{}

type id struct {
	idStart int
	idEnd   int
}

func (s *Solution) Day5part1() {
	idsSend, idsAvailable := getInput()
	idArray := make([]id, 0, len(idsAvailable))
	for _, line := range idsSend {
		parts := strings.SplitN(line, "-", 2)
		idArray = append(idArray, id{
			idStart: atoi(parts[0]),
			idEnd:   atoi(parts[1]),
		})
	}

	idAvailableArray := make([]int, 0, len(idsAvailable))
	for _, strId := range idsAvailable {
		idAvailableArray = append(idAvailableArray, atoi(strId))
	}

	sort.Slice(idAvailableArray, func(i, j int) bool {
		return idAvailableArray[i] < idAvailableArray[j]
	})

	sort.Slice(idArray, func(i, j int) bool {
		return idArray[i].idStart < idArray[j].idStart
	})

	counter := 0

	for _, normalId := range idAvailableArray {
		for _, idRange := range idArray {
			if normalId >= idRange.idStart && normalId <= idRange.idEnd {
				counter++
				break
			}

			if normalId < idRange.idStart {
				break
			}
		}
	}
	fmt.Println(counter)

}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getInput() ([]string, []string) {
	input, err := os.ReadFile("input/day5.txt")
	if err != nil {
		panic(err)
	}

	content := strings.ReplaceAll(string(input), "\r\n", "\n")
	split := strings.SplitN(content, "\n\n", 2)

	part1Lines := strings.Split(split[0], "\n")
	part2Lines := strings.Split(split[1], "\n")

	return part1Lines, part2Lines
}
