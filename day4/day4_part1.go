package day4

import (
	"fmt"
	"os"
	"strings"
)

type Solution struct{}

var lineLength int = 0

type point struct {
	x         int
	y         int
	character rune
}

func safeChar(points []point, width, x, y int) (rune, bool) {
	if width <= 0 {
		return 0, false
	}
	if x < 0 || y < 0 {
		return 0, false
	}
	idx := y*width + x
	if idx < 0 || idx >= len(points) {
		return 0, false
	}
	return points[idx].character, true
}

var points = getInput()

func (s *Solution) Day4part1() int {
	fmt.Println(lineLength)
	counter := 0

	for _, p := range points {
		if p.character == '@' {
			adiacencyPoints := 0
			if ch, ok := safeChar(points, lineLength, p.x-1, p.y); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x+1, p.y); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x, p.y-1); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x, p.y+1); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x-1, p.y-1); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x+1, p.y-1); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x-1, p.y+1); ok && ch == '@' {
				adiacencyPoints++
			}
			if ch, ok := safeChar(points, lineLength, p.x+1, p.y+1); ok && ch == '@' {
				adiacencyPoints++
			}
			if adiacencyPoints < 4 {
				counter++
				points[p.y*lineLength+p.x].character = 'X'
			}
		}
	}

	return counter
}

func getInput() []point {
	data, err := os.ReadFile("input/day4.txt")
	if err != nil {
		panic(err)
	}

	strData := string(data)
	splitLines := strings.Split(strData, "\n")
	lineLengthSet := false
	var points []point
	for i, strLine := range splitLines {

		strLine = strings.TrimSpace(strLine)

		if !lineLengthSet {
			lineLength = len(strLine)
			lineLengthSet = true
		}

		for j, letter := range strLine {
			points = append(points, point{x: j, y: i, character: letter})
		}
	}
	return points
}
