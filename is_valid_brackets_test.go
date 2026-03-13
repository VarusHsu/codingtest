package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidBrackets(t *testing.T) {
	assert.Equal(t, true, IsValidBrackets(""))
	assert.Equal(t, true, IsValidBrackets("()"))
	assert.Equal(t, false, IsValidBrackets("(]"))
	assert.Equal(t, false, IsValidBrackets("(}"))
	assert.Equal(t, true, IsValidBrackets("(())"))
	assert.Equal(t, true, IsValidBrackets("()()"))
	assert.Equal(t, false, IsValidBrackets("()()("))
	assert.Equal(t, true, IsValidBrackets("()[]()({})"))
	assert.Equal(t, false, IsValidBrackets("(-()()()()()()"))
}
