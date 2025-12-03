package day2

import "fmt"

func (s *Solution) Day2part2() {
	mapDistance := getInput()
	var sum int64

	for key, value := range mapDistance {
		for i := key; i <= value; i++ {
			if verifyNumberAtLeastTwo(i) {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}

func verifyNumberAtLeastTwo(num int64) bool {
	numStr := parseInt64ToString(num)
	n := len(numStr)
	for i := 2; i <= n; i++ {
		if n%i != 0 {
			continue
		}

		partLen := n / i
		firstPart := numStr[0:partLen]
		matched := true
		for j := 0; j < n; j += partLen {
			if numStr[j:j+partLen] != firstPart {
				matched = false
				break
			}
		}

		if matched {
			return true
		}
	}

	return false
}
