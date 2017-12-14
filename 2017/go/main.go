package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	_ "github.com/foophoof/advent-of-code/2017/go/day1"
	"github.com/foophoof/advent-of-code/2017/go/solution"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s path/to/input\n", os.Args[0])
		os.Exit(1)
	}
	inputDir := os.Args[1]

	for i := 1; i <= 25; i++ {
		input, err := ioutil.ReadFile(path.Join(inputDir, fmt.Sprintf("%02d.in", i)))
		if err != nil {
			fmt.Printf("Day %d: Couldn't read input: %v\n", i, err)
			continue
		}
		inputStr := strings.TrimSpace(string(input))
		outputPart1 := solution.CallPart1(i, inputStr)
		outputPart2 := solution.CallPart2(i, inputStr)
		fmt.Printf("Day %d:\n\tPart 1: %s\n\tPart 2: %s\n", i, outputPart1, outputPart2)
	}
}
