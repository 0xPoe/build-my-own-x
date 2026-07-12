package questions

import (
	"testing"

	"github.com/hi-rustin/lg/src/utils"
	"github.com/stretchr/testify/assert"
)

func TestMergeTwoLists(t *testing.T) {
	list1 := utils.BuildList([]int{1, 2, 3, 4, 5})
	list2 := utils.BuildList([]int{1, 2, 3, 4, 5})
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, utils.TraverseList(mergeTwoLists(list1, list2)))

	list1 = utils.BuildList([]int{2})
	list2 = utils.BuildList([]int{1})
	assert.Equal(t, []int{1, 2}, utils.TraverseList(mergeTwoLists(list1, list2)))
}
