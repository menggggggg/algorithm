package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TwoSum(t *testing.T) {
	array := []int{1, 3, 5, 6, 7, 8, 11, 23, 44, 55}

	tests := []struct {
		give int
		want bool
	}{
		{
			give: 4,
			want: true,
		},
		{
			give: 6,
			want: true,
		},
		{
			give: 12,
			want: true,
		},
		{
			give: 19,
			want: true,
		},
		{
			give: 2,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			ret := TwoSum(array, tt.give)
			assert.Equal(t, tt.want, ret)
		})
	}
}
