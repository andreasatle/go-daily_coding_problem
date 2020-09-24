package prereq_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/prereq"
	"github.com/stretchr/testify/assert"
)

func TestPrereq(t *testing.T) {
	// Non-cyclic case with result ["CS100", "CS200", "CS300"]
	m1 := prereq.StringSliceMap{"CS300": {"CS100", "CS200"}, "CS200": {"CS100"}, "CS100": {}}
	s1 := prereq.GetOrderedList(m1)
	assert.Equal(t, 3, len(s1), "Wrong lenght of output slice.")
	assert.Equal(t, "CS100", s1[0])
	assert.Equal(t, "CS200", s1[1])
	assert.Equal(t, "CS300", s1[2])

	// Cyclic case with result []
	m2 := prereq.StringSliceMap{"CS300": {"CS100", "CS200"}, "CS200": {"CS100", "CS300"}, "CS100": {}}
	s2 := prereq.GetOrderedList(m2)
	assert.Equal(t, 0, len(s2), "Wrong lenght of output slice.")
}
