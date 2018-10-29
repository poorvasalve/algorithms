package suffixarray

import (
	"testing"
	"gotest.tools/assert"
)

func TestSuffixArray(t *testing.T) {
	actual := buildSuffixArray("abbcbacba")
	expected := []suffix{
		{
			index:0,
			rank:[2]int{1:0},
		},
	}
	assert.Equal(t, expected, actual)
}