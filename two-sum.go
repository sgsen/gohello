package main

import "fmt"

func TwoSum1(nums []int, target int) []int {

	contains := func(slice []int, val int) (bool, int) {
		for idx, item := range slice {
			if item == val {
				return true, idx
			}
		}
		return false, -1
	}

	returnValue := []int{-1, -1}
	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Length of array: %d\n", len(nums))
	fmt.Printf("Target: %d\n", target)
	for index, value := range nums {
		fmt.Printf("index: %d ", index)
		fmt.Printf("value: %d\n", value)

		// see if the pair is available in the length of the array
		searchValue := target - value
		fmt.Printf("searchValue: %d ", searchValue)
		found, foundIndex := contains(nums[index+1:], searchValue)
		fmt.Printf("Item Found? %t ", found)
		if found {
			fmt.Printf("Item index? %v\n", foundIndex)
			return []int{index, index + foundIndex + 1}
		}
	}
	return returnValue
}

func TwoSum2(nums []int, target int) []int {
	// create a map to store values
	prevValues := make(map[int]int)

	// go through each item in the given array
	for curIndex, curValue := range nums {
		// what are you looking for
		searchValue := target - curValue

		// check map for searchValue
		searchValueIndex, exists := prevValues[searchValue]

		// if you're found it, return and exit
		if exists {
			return []int{searchValueIndex, curIndex}
		} else {
			// add this to the map
			prevValues[curValue] = curIndex
		}

	}
	fmt.Println("Map: ", prevValues)
	return []int{}
}

func test_two_sum() {
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
