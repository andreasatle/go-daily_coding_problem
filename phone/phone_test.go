package phone_test

import (
	"fmt"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/phone"
	"github.com/stretchr/testify/assert"
)

func TestPhoneCase1(t *testing.T) {
	strs := phone.Encode("23")
	fmt.Println(strs)
	assert.Equal(t, 9, len(strs))
	assert.Equal(t, "ad", strs[0])
	assert.Equal(t, "ae", strs[1])
	assert.Equal(t, "af", strs[2])
	assert.Equal(t, "bd", strs[3])
	assert.Equal(t, "be", strs[4])
	assert.Equal(t, "bf", strs[5])
	assert.Equal(t, "cd", strs[6])
	assert.Equal(t, "ce", strs[7])
	assert.Equal(t, "cf", strs[8])
}

func TestPhoneCase2(t *testing.T) {
	strs := phone.Encode("(832)488-9813")
	assert.Equal(t, 26244, len(strs))
	assert.Equal(t, "(tda)gtt-wt1d", strs[0])
	assert.Equal(t, "(tda)gtt-wt1e", strs[1])
	assert.Equal(t, "(tda)gtt-wt1f", strs[2])
	assert.Equal(t, "(tda)gtt-wu1d", strs[3])
}
