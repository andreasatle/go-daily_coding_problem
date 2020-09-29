package allperm_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/allperm"
	"github.com/stretchr/testify/assert"
)

func TestAllPerm(t *testing.T) {
	perms := allperm.Permutations([]int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, perms[0])
	assert.Equal(t, []int{1, 3, 2}, perms[1])
	assert.Equal(t, []int{2, 1, 3}, perms[2])
	assert.Equal(t, []int{2, 3, 1}, perms[3])
	assert.Equal(t, []int{3, 1, 2}, perms[4])
	assert.Equal(t, []int{3, 2, 1}, perms[5])
}
