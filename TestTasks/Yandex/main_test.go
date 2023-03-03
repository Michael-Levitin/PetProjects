package main

import (
	"fmt"
	"testing"
)

func TestCompareStrings(t *testing.T) {
	testCases := []struct {
		s1   string
		s2   string
		want int
	}{
		{"1.0", "1", 0},
		{"2.0.1", "2.0", 1},
		{"2.1.0", "2.10", -1},
		{"3.01", "3.1", 0},
		{"4.20.0-1", "4.020.00", 1},
		{"0.60-1", "0.60.0-1", 0},
		{"2a.0-1", "2.0", 2},
		{"7.0-alpha", "7.0.0-beta", -1},
		{"-1", "", 2},
	}

	for _, tc := range testCases {
		got := compareVersions(tc.s1, tc.s2)
		if got != tc.want {
			fmt.Printf("FAILED: for strings \"%s\" and \"%s\", got %d, want %d\n", tc.s1, tc.s2, got, tc.want)
		} else {
			fmt.Printf("PASSED: for strings \"%s\" and \"%s\", got %d\n", tc.s1, tc.s2, got)
		}
	}
}
