package day2_test

import (
	"testing"

	"github.com/foophoof/advent-of-code/2017/go/day2"
)

func TestChecksum_Part1(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{input: "5\t1\t9\t5\n7\t5\t3\n2\t4\t6\t8", output: "18"},
	}

	for i, testCase := range testCases {
		actualOutput := day2.ChecksumPart1(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("testCases[%d]: ChecksumPart1(%q) = %s, expected %s", i, testCase.input, actualOutput, testCase.output)
		}
	}
}

func TestChecksum_Part2(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{input: "5\t9\t2\t8\n9\t4\t7\t3\n3\t8\t6\t5", output: "9"},
	}

	for i, testCase := range testCases {
		actualOutput := day2.ChecksumPart2(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("testCases[%d]: ChecksumPart2(%q) = %s, expected %s", i, testCase.input, actualOutput, testCase.output)
		}
	}
}
