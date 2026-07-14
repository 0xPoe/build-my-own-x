package questions

import "github.com/hi-rustin/lg/src/utils"

func lowestCommonAncestor(root, p, q *utils.TreeNode) *utils.TreeNode {
	if root == nil {
		return nil
	}

	pVal := min(p.Val, q.Val)
	qVal := max(p.Val, q.Val)
	if root.Val >= pVal && root.Val <= qVal {
		return root
	} else if root.Val >= q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
