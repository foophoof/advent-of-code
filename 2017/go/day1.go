package main

import (
	"fmt"
	"os"
)

func solveCaptchaPart1(inputStr string) int {
	input := []byte(inputStr)
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+1)%len(input)] {
			sum += int(input[i] - '0')
		}
	}
	return sum
}

func solveCaptchaPart2(inputStr string) int {
	input := []byte(inputStr)
	sum := 0
	for i := 0; i < len(input); i++ {
		if input[i] == input[(i+len(input)/2)%len(input)] {
			sum += int(input[i] - '0')
		}
	}
	return sum
}

func main() {
	input := os.Args[1]
	fmt.Printf("Part 1: %d\n", solveCaptchaPart1(input))
	fmt.Printf("Part 2: %d\n", solveCaptchaPart2(input))
}
