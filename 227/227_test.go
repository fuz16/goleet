package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	assert.Equal(t, 5, calculate("3+2"))
	assert.Equal(t, 7, calculate("3+2*2"))
	assert.Equal(t, 5, calculate(" 3+5 / 2 "))
	assert.Equal(t, -1, calculate("1-2"))
}
