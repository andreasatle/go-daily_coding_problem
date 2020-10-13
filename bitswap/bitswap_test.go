package bitswap_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/bitswap"
	"github.com/stretchr/testify/assert"
)

func TestBitswap(t *testing.T) {
	// Test written with hexadecimal numbers, unclear how to
	// represent binary numbers in golang
	assert.Equal(t, uint8(0xAA), bitswap.Apply(uint8(0x55)))
	assert.Equal(t, uint8(0x55), bitswap.Apply(uint8(0xAA)))
	assert.Equal(t, uint8(0xE2), bitswap.Apply(uint8(0xD1)))
	assert.Equal(t, uint8(0xD1), bitswap.Apply(uint8(0xE2)))
}
