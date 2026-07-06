package questions

import "github.com/hi-rustin/lg/src/utils"

func reverseList(head *utils.ListNode) *utils.ListNode {
	if head == nil {
		return nil
	}
	var prev *utils.ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
