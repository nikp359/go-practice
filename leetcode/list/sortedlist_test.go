package list

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		list1    *ListNode
		list2    *ListNode
		wantList *ListNode
	}{
		{
			list1:    newListNode([]int{1, 3, 6}),
			list2:    newListNode([]int{2, 4}),
			wantList: newListNode([]int{1, 2, 3, 4, 6}),
		},
	}

	for i, tc := range testCases {
		t.Run(fmt.Sprintf("Case: %d", i), func(t *testing.T) {
			got := mergeTwoLists(tc.list1, tc.list2)

			assert.Equal(t, tc.wantList, got)
		})
	}
}

func newListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	ln := &ListNode{}
	base := ln

	for i, n := range nums {
		ln.Val = n

		if i+1 == len(nums) {
			return base
		}

		ln.Next = &ListNode{}
		ln = ln.Next
	}

	return base
}
