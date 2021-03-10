package leet

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canConstruct(t *testing.T) {
	assert.True(t, canConstruct("aa", "aab"))
	assert.False(t, canConstruct("aa", "ab"))
}
