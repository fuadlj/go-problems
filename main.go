package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val, " ")
		head = head.Next
	}
	fmt.Println()
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head
	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

func main() {
	list1 := createList([]int{1, 2, 4})
	list2 := createList([]int{1, 3, 4})
	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList)

	list1 = createList([]int{})
	list2 = createList([]int{})
	mergedList = mergeTwoLists(list1, list2)
	printList(mergedList)

	list1 = createList([]int{})
	list2 = createList([]int{0})
	mergedList = mergeTwoLists(list1, list2)
	printList(mergedList)
}
