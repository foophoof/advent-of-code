package day1_test

import (
	"testing"

	"github.com/foophoof/advent-of-code/2017/go/day1"
)

func TestSolveCaptcha_Part1(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{input: "1122", output: "3"},
		{input: "1111", output: "4"},
		{input: "1234", output: "0"},
		{input: "91212129", output: "9"},
	}

	for i, testCase := range testCases {
		actualOutput := day1.SolveCaptchaPart1(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("testCases[%d]: solveCaptcha(%q) = %s, expected %s", i, testCase.input, actualOutput, testCase.output)
		}
	}
}

func TestSolveCaptcha_Part2(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{input: "1212", output: "6"},
		{input: "1221", output: "0"},
		{input: "123425", output: "4"},
		{input: "123123", output: "12"},
		{input: "12131415", output: "4"},
	}

	for i, testCase := range testCases {
		actualOutput := day1.SolveCaptchaPart2(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("testCases[%d]: solveCaptcha(%q) = %s, expected %s", i, testCase.input, actualOutput, testCase.output)
		}
	}
}
