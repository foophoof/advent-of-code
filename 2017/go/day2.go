package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/foophoof/advent-of-code/2017/go/solution"
)

func init() {
	solution.RegisterPart1(2, ChecksumPart1)
	solution.RegisterPart2(2, ChecksumPart2)
}

func parseSpreadsheet(inputStr string) [][]int {
	var spreadsheet [][]int
	for _, line := range strings.Split(inputStr, "\n") {
		var row []int
		for _, cell := range strings.Split(line, "\t") {
			num, err := strconv.Atoi(cell)
			if err != nil {
				panic(fmt.Sprintf("couldn't parse number in spreadsheet: %v", err))
			}
			row = append(row, num)
		}
		spreadsheet = append(spreadsheet, row)
	}
	return spreadsheet
}

func ChecksumPart1(inputStr string) string {
	spreadsheet := parseSpreadsheet(inputStr)
	sum := 0
	for _, row := range spreadsheet {
		min, max := row[0], row[0]
		for _, cell := range row {
			if cell < min {
				min = cell
			}
			if cell > max {
				max = cell
			}
		}
		sum += max - min
	}
	return strconv.Itoa(sum)
}

func ChecksumPart2(inputStr string) string {
	spreadsheet := parseSpreadsheet(inputStr)
	sum := 0
	for _, row := range spreadsheet {
	rowLoop:
		for i := 0; i < len(row)-1; i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i]%row[j] == 0 {
					sum += row[i] / row[j]
					break rowLoop
				}
				if row[j]%row[i] == 0 {
					sum += row[j] / row[i]
					break rowLoop
				}
			}
		}
	}
	return strconv.Itoa(sum)
}
