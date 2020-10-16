package pi_test

import (
	"math"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/pi"
	"github.com/stretchr/testify/assert"
)

func TestPi(t *testing.T) {
	assert.True(t, math.Abs(pi.Approx()-math.Pi) < 0.0005)
}

func TestPiBatch(t *testing.T) {
	assert.True(t, math.Abs(pi.ApproxBatch(1000)-math.Pi) < 0.0005)
}
