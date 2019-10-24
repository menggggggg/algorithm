package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Reverse(t *testing.T) {
	n5 := New("5", nil)
	n4 := New("4", n5)
	n3 := New("3", n4)
	n2 := New("2", n3)
	n1 := New("1", n2)
	l := Reverse(n1)
	assert.Equal(t, "5", l.Data())
	l = l.Next()
	assert.Equal(t, "4", l.Data())
	l = l.Next()
	assert.Equal(t, "3", l.Data())
	l = l.Next()
	assert.Equal(t, "2", l.Data())
	l = l.Next()
	assert.Equal(t, "1", l.Data())
	l = l.Next()
	assert.Nil(t, l)
}
