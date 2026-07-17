package questions

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)

	for _, group := range result {
		sort.Strings(group)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	expected := [][]string{
		{"ate", "eat", "tea"},
		{"bat"},
		{"nat", "tan"},
	}
	assert.Equal(t, expected, result)
}
