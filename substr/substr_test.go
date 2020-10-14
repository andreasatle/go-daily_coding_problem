package substr_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/substr"
	"github.com/stretchr/testify/assert"
)

func TestSubStr1(t *testing.T) {
	byteSet := substr.NewByteSet("abcd")
	subStr, ok := byteSet.SubSet("hbgcafdhpqwrgbc")
	assert.Equal(t, "bgcafd", subStr)
	assert.True(t, ok)
}
func TestSubStr2(t *testing.T) {
	byteSet := substr.NewByteSet("aei")
	subStr, ok := byteSet.SubSet("figehaeci")
	assert.Equal(t, "aeci", subStr)
	assert.True(t, ok)
}

func TestSubStr3(t *testing.T) {
	byteSet := substr.NewByteSet("aqi")
	subStr, ok := byteSet.SubSet("figehaeci")
	assert.Equal(t, "", subStr)
	assert.False(t, ok)
}
