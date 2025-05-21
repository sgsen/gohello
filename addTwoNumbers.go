package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddToLList(firstNode *ListNode, value int) *ListNode {
	// create a node holding the value
	newNode := ListNode{value, nil}

	// check if the firstNode is already a list or if it's empty
	if firstNode == nil {
		return &newNode
	}

	// point to the given node and iterate until you get to node pointing to nil which is the last node
	current := firstNode
	for current.Next != nil {
		current = current.Next
	}
	// when you exit this while loop, you are pointing to the last node with current
	current.Next = &newNode
	return firstNode
}

func PrintLList(firstNode *ListNode) {
	fmt.Println("\nPrinting linked list...")
	current := firstNode
	for current.Next != nil {
		fmt.Println(current.Val)
		current = current.Next
	}
	// the value of the last list item
	fmt.Println(current.Val)
}

func CreateLLfromList(intList []int) *ListNode {
	var newLinkedList *ListNode = nil
	for _, value := range intList {
		newLinkedList = AddToLList(newLinkedList, value)
	}
	return newLinkedList
}

func ReverseLList(firstNode *ListNode) *ListNode {
	currNode := firstNode
	prevNode := (*ListNode)(nil)

	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode

	}
	return prevNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sumLL *ListNode = nil
	var val1, val2, sumVal, digit int
	var carry int = 0

	current1 := l1
	current2 := l2

	for (current1 != nil) || (current2 != nil) {
		val1 = 0
		val2 = 0

		if current1 != nil {
			val1 = current1.Val
			current1 = current1.Next
		}

		if current2 != nil {
			val2 = current2.Val
			current2 = current2.Next
		}

		sumVal = val1 + val2 + carry
		carry = sumVal / 10
		digit = sumVal % 10

		// fmt.Printf("val1: %v, val2: %v, sumVal: %v, digit: %v, carry: %v\n", val1, val2, sumVal, digit, carry)

		sumLL = AddToLList(sumLL, digit)

	}
	// deal with the case where there is a carry left over
	if carry > 0 {
		sumLL = AddToLList(sumLL, carry)
	}

	return sumLL
}
