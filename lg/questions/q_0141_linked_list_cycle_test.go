package questions

import (
	"testing"

	"github.com/hi-rustin/lg/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestHasCycle(t *testing.T) {
	t.Run("cycle", func(t *testing.T) {
		node1 := &utils.ListNode{Val: 3}
		node2 := &utils.ListNode{Val: 2}
		node3 := &utils.ListNode{Val: 0}
		node4 := &utils.ListNode{Val: -4}
		node1.Next = node2
		node2.Next = node3
		node3.Next = node4
		node4.Next = node2

		assert.True(t, hasCycle(node1))
	})

	t.Run("self cycle", func(t *testing.T) {
		head := &utils.ListNode{Val: 1}
		head.Next = head

		assert.True(t, hasCycle(head))
	})

	t.Run("no cycle", func(t *testing.T) {
		head := utils.BuildList([]int{1, 2, 3})

		assert.False(t, hasCycle(head))
	})

	t.Run("empty list", func(t *testing.T) {
		assert.False(t, hasCycle(nil))
	})
}
