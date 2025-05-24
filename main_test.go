package main

import (
	"testing"
)

type TestCase struct {
	input  string
	output int
}

func TestAddTwoNumbers(t *testing.T) {
	// fmt.Println(quote.Go())
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}

	var ll1 *ListNode = CreateLLfromList(l1)
	var ll2 *ListNode = CreateLLfromList(l2)
	var sumLL *ListNode = nil

	PrintLList(ll1)
	PrintLList(ll2)
	sumLL = AddTwoNumbers(ll1, ll2)
	// revll1 := ReverseLList(ll1)
	PrintLList(sumLL)
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testCases := []TestCase{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"aab", 2},
		{"aabaab!bb", 3},
	}

	for _, tc := range testCases {
		result := lengthOfLongestSubstring(tc.input)
		expected := tc.output

		if result != expected {
			t.Errorf("Failed %s. Expected %d, but got %d", tc.input, expected, result)
		} else {
			println("Passed: ", tc.input)
		}
	}
}
