package main

func main() {
	// var l1 = []int{2, 4, 3}
	// var l2 = []int{5, 6, 4}

	// var l1 = []int{0}
	// var l2 = []int{0}

	var l1 = []int{9, 9, 9, 9, 9, 9, 9}
	var l2 = []int{9, 9, 9, 9}

	var ll1 *ListNode = CreateLLfromList(l1)
	var ll2 *ListNode = CreateLLfromList(l2)
	var sumLL *ListNode = nil

	PrintLList(ll1)
	PrintLList(ll2)
	sumLL = addTwoNumbers(ll1, ll2)
	//revll1 := ReverseLList(ll1)
	PrintLList(sumLL)

}
