package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceBased(t *testing.T) {
	s := NewSliceBased()
	if s == nil {
		t.Error("NewSliceBased() returned nil")
	}

	assert.Equal(t, true, s.IsEmpty())
	i, ok := s.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, i)

	s.Push(1)
	assert.Equal(t, false, s.IsEmpty())
	i, ok = s.Top()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, i)

	i, ok = s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, i)
	assert.Equal(t, true, s.IsEmpty())
}
