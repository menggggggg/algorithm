package stringint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StringToInt(t *testing.T) {
	tests := []struct {
		give string
		want int
	}{
		{
			give: "123456",
			want: 123456,
		},
		{
			give: "-123456",
			want: -123456,
		},
	}

	for _, tt := range tests {
		t.Run(tt.give, func(t *testing.T) {
			i, err := StringToInt(tt.give)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, i)
		})
	}
}

func Test_IntToString(t *testing.T) {
	tests := []struct {
		give int
		want string
	}{
		{
			give: 123456,
			want: "123456",
		},
		{
			give: -123456,
			want: "-123456",
		},
	}

	for _, tt := range tests {
		t.Run(tt.want, func(t *testing.T) {
			s := IntToString(tt.give)
			assert.Equal(t, tt.want, s)
		})
	}
}
