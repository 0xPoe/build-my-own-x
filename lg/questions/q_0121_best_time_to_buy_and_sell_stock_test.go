package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	testCase := []int{7, 1, 5, 3, 6, 4}
	assert.Equal(t, 5, maxProfit(testCase))
	testCase = []int{7, 6, 4, 3, 1}
	assert.Equal(t, 0, maxProfit(testCase))
}
