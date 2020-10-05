package substr_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/substr"
	"github.com/stretchr/testify/assert"
)

func TestSubStr1(t *testing.T) {
	byteSet := substr.NewByteSet("abcd")
	subStr, _ := byteSet.SubSet("hbgcafdhpqwrgbc")
	assert.Equal(t, "bgcafd", subStr)
}
func TestSubStr2(t *testing.T) {
	byteSet := substr.NewByteSet("aei")
	subStr, _ := byteSet.SubSet("figehaeci")
	assert.Equal(t, "aeci", subStr)
}
