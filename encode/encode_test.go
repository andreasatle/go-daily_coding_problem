package encode_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/encode"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	encMsg, _ := encode.Encode("bgbg")
	assert.Equal(t, 1, encode.DecodeCount(encMsg))

	encMsg, _ = encode.Encode("bfbg")
	assert.Equal(t, 2, encode.DecodeCount(encMsg))

	encMsg, _ = encode.Encode("abc")
	assert.Equal(t, 3, encode.DecodeCount(encMsg))

	encMsg, _ = encode.Encode("foobar")
	assert.Equal(t, 20, encode.DecodeCount(encMsg))

	// Encode an illegal string
	_, err := encode.Encode("foo-bar")
	assert.False(t, err == nil)

	// Decode an illegal string
	assert.Equal(t, 0, encode.DecodeCount("001"))

	// Decode an illegal string
	assert.Equal(t, 0, encode.DecodeCount("100"))
}
