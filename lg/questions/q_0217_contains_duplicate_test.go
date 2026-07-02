package questions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	cases := []struct {
		nums     []int
		expected bool
	}{
		{
			nums:     []int{1, 2, 3, 3},
			expected: true,
		},
		{
			nums:     []int{1, 2, 3, 4},
			expected: false,
		},
	}
	for _, testCase := range cases {

		assert.Equal(t, testCase.expected, containsDuplicate(testCase.nums))
	}
}
