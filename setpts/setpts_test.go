package setpts

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetpts(t *testing.T) {
	iVals := Intervals{
		Interval{0, 3},
		Interval{2, 6},
		Interval{3, 4},
		Interval{6, 9},
	}
	points := iVals.FindPoints()
	assert.Equal(t, 2, len(points))
	assert.Equal(t, 3, points[0])
	assert.Equal(t, 6, points[1])
}
