package questions

import (
	"testing"

	"github.com/hi-rustin/lg/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestInvertBinaryTree(t *testing.T) {
	tree := utils.BuildTree([]int{4, 2, 7, 1, 3, 6, 9})
	assert.Equal(t, []int{4, 7, 2, 9, 6, 3, 1}, utils.TraverseTree(invertTree(tree)))
}
