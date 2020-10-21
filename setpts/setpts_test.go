package setpts_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/setpts"
	"github.com/stretchr/testify/assert"
)

func TestSetpts(t *testing.T) {
	iVals := setpts.Intervals{
		setpts.Interval{0, 3},
		setpts.Interval{2, 6},
		setpts.Interval{3, 4},
		setpts.Interval{6, 9},
	}
	points := iVals.FindPoints()
	assert.Equal(t, 2, len(points))
	assert.Equal(t, 3, points[0])
	assert.Equal(t, 6, points[1])
}
