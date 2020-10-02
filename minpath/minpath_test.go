package minpath_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/minpath"
	"github.com/stretchr/testify/assert"
)

func TestMinPath(t *testing.T) {
	assert.Equal(t, 2, minpath.MinSteps([]minpath.Point{
		minpath.Point{0, 0},
		minpath.Point{1, 1},
		minpath.Point{1, 2},
	}))

	assert.Equal(t, 12, minpath.MinSteps([]minpath.Point{
		minpath.Point{0, 0},
		minpath.Point{3, 2},
		minpath.Point{2, -7},
	}))
}
