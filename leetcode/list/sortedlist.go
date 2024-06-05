package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	base := &ListNode{}

	merge(list1, list2, base)

	return base
}

func merge(list1 *ListNode, list2 *ListNode, merger *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	if list1 == nil {
		merger = list2
		return nil
	}

	if list2 == nil {
		merger = list1
		return nil
	}

	val := list1.Val

	if list1.Val >= list2.Val {
		val = list2.Val
	}

	merger.Val = val
	merger.Next = &ListNode{}
	return merge(list1, list2.Next, merger.Next)
}
