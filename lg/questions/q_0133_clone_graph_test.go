package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloneGraph(t *testing.T) {
	// adjList = [[2,4],[1,3],[2,4],[1,3]]
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node1.Neighbors = []*Node{node2, node4}
	node2.Neighbors = []*Node{node1, node3}
	node3.Neighbors = []*Node{node2, node4}
	node4.Neighbors = []*Node{node1, node3}

	clone := cloneGraph(node1)

	assert.NotSame(t, node1, clone)
	assert.Equal(t, 1, clone.Val)
	assert.Len(t, clone.Neighbors, 2)
	assert.Equal(t, 2, clone.Neighbors[0].Val)
	assert.Equal(t, 4, clone.Neighbors[1].Val)
	// following the edge 1 -> 2 -> 1 must lead back to the same cloned node
	assert.Same(t, clone, clone.Neighbors[0].Neighbors[0])

	// adjList = [[]]
	single := &Node{Val: 1}
	cloneSingle := cloneGraph(single)
	assert.NotSame(t, single, cloneSingle)
	assert.Equal(t, 1, cloneSingle.Val)
	assert.Empty(t, cloneSingle.Neighbors)

	// adjList = []
	assert.Nil(t, cloneGraph(nil))
}
