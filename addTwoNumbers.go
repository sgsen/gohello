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

func PrintList(firstNode *ListNode) {
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	PrintList(l1)
	PrintList(l2)
	return l1
}
