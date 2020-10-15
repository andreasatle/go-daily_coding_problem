package rooms_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/rooms"
	"github.com/stretchr/testify/assert"
)

func TestRooms(t *testing.T) {
	ivals := rooms.Intervals{rooms.Interval{30, 75}, rooms.Interval{0, 50}, rooms.Interval{60, 150}}
	assert.Equal(t, 2, rooms.Minimum(ivals))

	ivals = append(ivals, rooms.Interval{30, 60})
	assert.Equal(t, 3, rooms.Minimum(ivals))
}
