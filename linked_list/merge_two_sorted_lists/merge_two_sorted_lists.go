package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	mergedHead := &ListNode{}
	var current *ListNode

	for list1 != nil && list2 != nil {
		if current == nil {
			current = mergedHead
		} else {
			current.Next = &ListNode{}
			current = current.Next
		}

		switch {
		case list1.Val < list2.Val:
			current.Val = list1.Val
			list1 = list1.Next
		case list2.Val < list1.Val:
			current.Val = list2.Val
			list2 = list2.Next
		case list1.Val == list2.Val:
			current.Val = list1.Val
			current.Next = &ListNode{list2.Val, nil}
			current = current.Next
			list1 = list1.Next
			list2 = list2.Next
		default:
			panic("uncomparable listNode values")
		}
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return mergedHead
}
