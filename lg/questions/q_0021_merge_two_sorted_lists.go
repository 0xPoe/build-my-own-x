package questions

import "github.com/hi-rustin/lg/src/utils"

func mergeTwoLists(list1 *utils.ListNode, list2 *utils.ListNode) *utils.ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var newHead, previous, current *utils.ListNode
	for list1 != nil && list2 != nil {
		var tmp *utils.ListNode
		if list1.Val <= list2.Val {
			tmp = list1
			list1 = list1.Next
		} else {
			tmp = list2
			list2 = list2.Next
		}

		previous = current
		current = tmp
		if previous != nil {
			previous.Next = current
			if newHead == nil {
				newHead = previous
			}
		}
	}
	if newHead == nil {
		newHead = current
	}

	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}

	return newHead
}
