package regexp_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/regexp"
	"github.com/stretchr/testify/assert"
)

func TestRegExp(t *testing.T) {
	assert.Equal(t, true, regexp.Check("My happy day", "* h..py*y"))
	assert.Equal(t, true, regexp.Check("", ""))
	assert.Equal(t, true, regexp.Check("happy", "....."))
	assert.Equal(t, true, regexp.Check("xyxyxyxyxyxy", "xy*x.x..*y"))
	assert.Equal(t, true, regexp.Check("*.*", ".*."))
	assert.Equal(t, false, regexp.Check("*.*", ".*b"))
	assert.Equal(t, true, regexp.Check("ray", "ray"))
	assert.Equal(t, true, regexp.Check("ray", "ra."))
	assert.Equal(t, false, regexp.Check("rayban", "ra."))
}
