package questions

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloned := make(map[*Node]*Node)
	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		newNode := &Node{
			Val:       n.Val,
			Neighbors: make([]*Node, len(n.Neighbors)),
		}
		if neighbor, ok := cloned[n]; ok {
			return neighbor
		}

		cloned[n] = newNode
		for i, neighbor := range n.Neighbors {
			newNode.Neighbors[i] = dfs(neighbor)
		}
		return newNode
	}
	return dfs(node)
}
