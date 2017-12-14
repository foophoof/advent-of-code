package solution

import (
	"fmt"
	"sync"
)

type Solution func(input string) string

var solutionsSetup sync.Once
var solutionsPart1 []Solution
var solutionsPart2 []Solution

func RegisterPart1(day int, f Solution) {
	solutionsSetup.Do(func() {
		solutionsPart1 = make([]Solution, 25)
		solutionsPart2 = make([]Solution, 25)
	})
	if day < 0 || day > 25 {
		panic(fmt.Sprintf("solution.Register got day %d, must be in range [1, 25]", day))
	}

	solutionsPart1[day-1] = f
}
func RegisterPart2(day int, f Solution) {
	solutionsSetup.Do(func() {
		solutionsPart1 = make([]Solution, 25)
		solutionsPart2 = make([]Solution, 25)
	})
	if day < 0 || day > 25 {
		panic(fmt.Sprintf("solution.Register got day %d, must be in range [1, 25]", day))
	}

	solutionsPart2[day-1] = f
}

func CallPart1(day int, input string) string {
	f := solutionsPart1[day-1]
	if f == nil {
		return ""
	}
	return f(input)
}

func CallPart2(day int, input string) string {
	f := solutionsPart2[day-1]
	if f == nil {
		return ""
	}
	return f(input)
}
