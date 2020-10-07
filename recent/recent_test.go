package recent_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/recent"
	"github.com/stretchr/testify/assert"
)

func TestRecent(t *testing.T) {

	m := recent.NewMap()

	m.Set(1, 1, 0)
	m.Set(1, 2, 2)

	val, err := m.Get(1, 1)
	assert.Equal(t, 1, val)
	assert.Equal(t, nil, err)

	val, err = m.Get(1, 3)
	assert.Equal(t, 2, val)
	assert.Equal(t, nil, err)

	m.Set(1, 1, 5)
	val, err = m.Get(1, -1)
	assert.NotEqual(t, nil, err)

	val, err = m.Get(1, 10)
	assert.Equal(t, 1, val)
	assert.Equal(t, nil, err)

	m.Set(1, 1, 0)
	m.Set(1, 2, 0)
	val, err = m.Get(1, 0)
	assert.Equal(t, 2, val)
	assert.Equal(t, nil, err)
}
