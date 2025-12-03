package day3

import (
	"fmt"
	"strings"
)

func (s *Solution) Day3part2() {
	banks := getInput()
	var sum int64
	k := 12

	for _, bank := range banks {
		bank = strings.TrimSpace(bank)
		digits := make([]byte, 0, len(bank))
		for i := 0; i < len(bank); i++ {
			if bank[i] >= '0' && bank[i] <= '9' {
				digits = append(digits, bank[i])
			}
		}

		if len(digits) < k {
			continue
		}

		val := pickMaxKDigits(digits, k)
		sum += val
	}

	fmt.Println(sum)
}

func pickMaxKDigits(digits []byte, k int) int64 {
	n := len(digits)
	toRemove := n - k
	stack := make([]byte, 0, k)

	for i := 0; i < n; i++ {
		digit := digits[i]
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}
		stack = append(stack, digit)
	}

	stack = stack[:k]

	var val int64
	for i := 0; i < k; i++ {
		val = val*10 + int64(stack[i]-'0')
	}
	return val
}
