package questions

import (
	"testing"

	"github.com/hi-rustin/lg/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {
	list := utils.BuildList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{5, 4, 3, 2, 1}, utils.TraverseList(reverseList(list)))
}
