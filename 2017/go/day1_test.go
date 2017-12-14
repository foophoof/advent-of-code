package main

import "testing"

func TestSolveCaptcha_Part1(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{input: "1122", output: 3},
		{input: "1111", output: 4},
		{input: "1234", output: 0},
		{input: "91212129", output: 9},
	}

	for i, testCase := range testCases {
		actualOutput := solveCaptchaPart1(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("testCases[%d]: solveCaptcha(%q) = %d, expected %d", i, testCase.input, actualOutput, testCase.output)
		}
	}
}

func TestSolveCaptcha_Part2(t *testing.T) {
	testCases := []struct {
		input  string
		output int
	}{
		{input: "1212", output: 6},
		{input: "1221", output: 0},
		{input: "123425", output: 4},
		{input: "123123", output: 12},
		{input: "12131415", output: 4},
	}

	for i, testCase := range testCases {
		actualOutput := solveCaptchaPart2(testCase.input)
		if actualOutput != testCase.output {
			t.Errorf("testCases[%d]: solveCaptcha(%q) = %d, expected %d", i, testCase.input, actualOutput, testCase.output)
		}
	}
}
