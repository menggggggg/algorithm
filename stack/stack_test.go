package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	st := New()

	st.Push(5)
	assert.Equal(t, 5, st.GetValue())

	st.Push(4)
	assert.Equal(t, 4, st.GetValue())

	st.Push(3)
	assert.Equal(t, 3, st.GetValue())

	st.Push(2)
	assert.Equal(t, 2, st.GetValue())

	st.Push(1)
	assert.Equal(t, 1, st.GetValue())

	assert.Equal(t, 5, st.Len())

	assert.Equal(t, 1, st.Pop())
	assert.Equal(t, 2, st.Pop())
	assert.Equal(t, 3, st.Pop())
	assert.Equal(t, 4, st.Pop())
	assert.Equal(t, 5, st.Pop())
	assert.Equal(t, nil, st.Pop())
}
