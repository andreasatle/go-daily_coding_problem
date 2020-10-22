package singleton_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/singleton"
	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	var data *singleton.Singleton
	getInstance := singleton.New()

	data = getInstance()
	assert.Equal(t, "This is odd!", data.SomeFunData)

	data = getInstance()
	assert.Equal(t, "Time to get even!", data.SomeFunData)

	data = getInstance()
	assert.Equal(t, "This is odd!", data.SomeFunData)

	data = getInstance()
	assert.Equal(t, "Time to get even!", data.SomeFunData)

	data = getInstance()
	assert.Equal(t, "This is odd!", data.SomeFunData)

	data = getInstance()
	assert.Equal(t, "Time to get even!", data.SomeFunData)
}
