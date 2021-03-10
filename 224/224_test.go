package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	assert.Equal(t, calculate("1 + 1"), 2)
	assert.Equal(t, calculate("(1 + 1)"), 2)
	assert.Equal(t, calculate("(1+(4+5+2)-3)+(6+8)"), 23)
	assert.Equal(t, calculate(" 2-1 + 2 "), 3)
	assert.Equal(t, calculate("-(3+(4+5))"), -12)
	assert.Equal(t, calculate("1-(5)"), -4)
}
