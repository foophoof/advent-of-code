package day1

import (
	"strconv"

	"github.com/foophoof/advent-of-code/2017/go/solution"
)

func init() {
	solution.RegisterPart1(1, SolveCaptchaPart1)
	solution.RegisterPart2(1, SolveCaptchaPart2)
}

func SolveCaptchaPart1(inputStr string) string {
	input := []byte(inputStr)
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			sum += int(input[i] - '0')
		}
	}
	return strconv.Itoa(sum)
}

func SolveCaptchaPart2(inputStr string) string {
	input := []byte(inputStr)
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+len(input)/2)%len(input)] {
			sum += int(input[i] - '0')
		}
	}
	return strconv.Itoa(sum)
}
