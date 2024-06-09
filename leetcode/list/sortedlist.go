package list

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

	var result, tmp *ListNode

	if list1.Val < list2.Val {
		tmp = &ListNode{Val: list1.Val}
		list1 = list1.Next
	} else {
		tmp = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}

	result = tmp

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tmp.Next = &ListNode{Val: list1.Val}
			list1 = list1.Next
		} else {
			tmp.Next = &ListNode{Val: list2.Val}
			list2 = list2.Next
		}

		tmp = tmp.Next
	}

	if list1 != nil {
		tmp.Next = list1
	}

	if list2 != nil {
		tmp.Next = list2
	}

	return result
}
