package sumk_test

import (
	"testing"

	"github.com/andreasatle/Assorted/Daily/sumk"
	"github.com/stretchr/testify/assert"
)

func TestSumK(t *testing.T) {
	assert.Equal(t, true, sumk.SumK([]int{10, 15, 3, 7}, 17), "10 and 7 add to 17.")
	assert.Equal(t, false, sumk.SumK([]int{10, 15, 3, 7}, 16), "No two numbers add to 16.")
}
