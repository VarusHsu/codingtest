package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	assert.Equal(t,
		MergeIntervals(
			[][]int{
				{1, 3},
				{2, 6},
				{8, 10},
				{15, 18},
			},
		),
		[][]int{
			{1, 6},
			{8, 10},
			{15, 18},
		},
	)
}
