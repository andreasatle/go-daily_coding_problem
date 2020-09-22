package rgb_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/rgb"
	"github.com/stretchr/testify/assert"
)

func TestRgb(t *testing.T) {
	rgbSlice := []rune{'B', 'R', 'G', 'R', 'R', 'B', 'B', 'G'}
	rgb.Sort(rgbSlice)
	assert.Equal(t, 'R', rgbSlice[0])
	assert.Equal(t, 'R', rgbSlice[1])
	assert.Equal(t, 'R', rgbSlice[2])
	assert.Equal(t, 'G', rgbSlice[3])
	assert.Equal(t, 'G', rgbSlice[4])
	assert.Equal(t, 'B', rgbSlice[5])
	assert.Equal(t, 'B', rgbSlice[6])
	assert.Equal(t, 'B', rgbSlice[7])
}
