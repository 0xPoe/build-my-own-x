package questions

import (
	"testing"

	"github.com/hi-rustin/lg/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	tree := utils.BuildTree([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5})
	assert.Equal(t, tree, lowestCommonAncestor(tree, tree.Left, tree.Right))
}
