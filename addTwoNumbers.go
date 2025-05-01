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

/* func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// get the last item from list1 and update what the last item is for next time
	l1
	currentl2 := l2

	return l1
}
*/
