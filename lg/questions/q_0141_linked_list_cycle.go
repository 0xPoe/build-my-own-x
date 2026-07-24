package questions

import "github.com/hi-rustin/lg/src/utils"

func hasCycle(head *utils.ListNode) bool {
	if head == nil {
		return false
	}
	nodeMap := make(map[*utils.ListNode]struct{})
	for head != nil {
		if _, ok := nodeMap[head]; ok {
			return true
		}
		nodeMap[head] = struct{}{}
		head = head.Next
	}
	return false
}
