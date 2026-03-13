package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindNumOnlyAppearOnce(t *testing.T) {
	assert.Equal(t, FindNumOnlyAppearOnce([]int{1, 1, 3, 4, 4}), 3)
	assert.Equal(t, FindNumOnlyAppearOnce([]int{1, 2, 3, 2, 1}), 3)
	assert.Equal(t, FindNumOnlyAppearOnce([]int{1, -1, -1, 4, 4, 9, 1}), 9)
	assert.Equal(t, FindNumOnlyAppearOnce([]int{1, 1, 3}), 3)
}
