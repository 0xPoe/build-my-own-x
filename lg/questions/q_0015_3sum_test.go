package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum(nums))
}
