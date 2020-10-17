package partition_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/partition"
	"github.com/stretchr/testify/assert"
)

func TestPartitionCase1(t *testing.T) {
	// Case 1 from problem formulation
	assert.True(t, partition.SameSum([]int{15, 5, 20, 10, 35, 15, 10}))
}

func TestPartitionCase2(t *testing.T) {
	// Case 2 from problem formulation
	assert.False(t, partition.SameSum([]int{15, 5, 20, 10, 35}))
}

func TestPartitionCase3(t *testing.T) {
	// Simple case which work
	assert.True(t, partition.SameSum([]int{3, 2, 4, 1}))
}

func TestPartitionCase4(t *testing.T) {
	// Total sum is odd
	assert.False(t, partition.SameSum([]int{3, 4, 1, 2, 7}))
}

func TestPartitionCase5(t *testing.T) {
	// Empty set, assuming sum({}) == 0
	assert.True(t, partition.SameSum([]int{}))
}
