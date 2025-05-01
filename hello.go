package main

import "fmt"

func main_two_sum() {
	//fmt.Println(quote.Go())

	// test case 2
	inputArray := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("Target Match: %v\n", TwoSum2(inputArray, target))

	// test case 2
	inputArray = []int{3, 2, 4}
	target = 6
	fmt.Printf("Target Match: %v\n", TwoSum2(inputArray, target))

	// test case 3
	inputArray = []int{3, 3}
	target = 6
	fmt.Printf("Target Match: %v\n", TwoSum2(inputArray, target))

}

func main() {
	var l1 = []int{1, 2, 3}
	//var l2 = []int{4, 5, 6}

	var ll1 *ListNode = CreateLLfromList(l1)
	//var ll2 *ListNode = CreateLLfromList(l2)

	PrintLList(ll1)
	// addTwoNumbers(ll1, ll2)
	revll1 := ReverseLList(ll1)
	PrintLList(revll1)

}
