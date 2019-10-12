package addtwo

import (
	"math"
)

//ListNode reverse order and each of their nodes contain a single digit
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return intToNode(nodeToInt(l1) + nodeToInt(l2))
}

func nodeToInt(l *ListNode) int {
	return rn(l, 0, 0)
}

func intToNode(num int) *ListNode {
	var rightInt int
	tmpInt := num

	node := &ListNode{}
	currentNode := node

	for {
		rightInt = tmpInt % 10
		currentNode.Val = rightInt

		if tmpInt /= 10; tmpInt == 0 {
			break
		}

		currentNode.Next = &ListNode{}
		currentNode = currentNode.Next
	}

	return node
}

func rn(l *ListNode, step int, accum int) int {
	accum = accum + l.Val*int(math.Pow10(step))
	step++

	if l.Next != nil {
		return rn(l.Next, step, accum)
	}

	return accum
}
