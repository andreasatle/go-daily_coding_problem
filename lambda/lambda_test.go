package lambda_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/lambda"
	"github.com/stretchr/testify/assert"
)

func TestLambda(t *testing.T) {
	get, acc, set, lessThan := lambda.CreateNext(0)

	assert.Equal(t, true, lessThan(3))
	assert.Equal(t, 0, get())
	acc()
	assert.Equal(t, true, lessThan(3))
	assert.Equal(t, 1, get())
	acc()
	assert.Equal(t, true, lessThan(3))
	assert.Equal(t, 2, get())
	acc()
	assert.Equal(t, false, lessThan(3))

	// Check multiply uses of get()
	// I choosed to have separate functions get and acc.
	set(4711)
	assert.Equal(t, 4711, get())
	assert.Equal(t, 4711, get())
}
