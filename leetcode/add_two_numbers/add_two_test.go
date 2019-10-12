package addtwo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	t.Run("node to int", func(t *testing.T) {
		ln := node465()
		got := nodeToInt(ln)
		assert.Equal(t, 465, got)
	})

	t.Run("int to node", func(t *testing.T) {
		node := intToNode(465)
		assert.Equal(t, 5, node.Val)
		assert.Equal(t, 6, node.Next.Val)
		assert.Equal(t, 4, node.Next.Next.Val)
	})

	t.Run("Add Two Numbers", func(t *testing.T) {
		got := addTwoNumbers(node465(), node301())
		assert.Equal(t, 6, got.Val)
		assert.Equal(t, 6, got.Next.Val)
		assert.Equal(t, 7, got.Next.Next.Val)
		assert.Equal(t, nil, got.Next.Next.Next)
	})

}

func node465() *ListNode {
	return &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
}

func node301() *ListNode {
	return &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
}

// Input: [1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
// [5,6,4]
// Output: [-2,-4,-3,-5,-7,-7,-4,-5,-8,-6,-3,0,-2,-7,-3,-3,-2,-2,-9]
// Expected: [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
