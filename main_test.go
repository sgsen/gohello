package main

import (
	"fmt"
	"testing"
)

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
	fmt.Println("running test_lengthOfLongestSubstring...")
	test_str := "hello"
	str_length := lengthOfLongestSubstring(test_str)
	fmt.Println("string length: ", str_length)
}
