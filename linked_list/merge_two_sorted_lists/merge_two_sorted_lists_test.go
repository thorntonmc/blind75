package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func newListFromVals(vals ...int) *ListNode {
	head := &ListNode{vals[0], nil}
	current := head
	vals = vals[1:]
	for _, val := range vals {
		l := &ListNode{val, nil}
		current.Next = l
		current = l
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{
			list1: newListFromVals(1, 2, 4),
			list2: newListFromVals(1, 3, 4),
			want:  newListFromVals(1, 1, 2, 3, 4, 4),
		}, {
			list1: nil,
			list2: nil,
			want:  nil,
		}, {
			list1: nil,
			list2: newListFromVals(0),
			want:  newListFromVals(0),
		}, {
			list1: newListFromVals(1, 2, 3, 4, 5),
			list2: newListFromVals(1, 2),
			want:  newListFromVals(1, 1, 2, 2, 3, 4, 5),
		},
	}

	for _, test := range tests {
		got := mergeTwoLists(test.list1, test.list2)
		assert.Equal(t, test.want, got)
	}
}
