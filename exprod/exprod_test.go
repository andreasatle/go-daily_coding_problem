package exprod_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/exprod"
	"github.com/stretchr/testify/assert"
)

func TestExProd1(t *testing.T) {
	prod := exprod.Product([]int{1, 2, 3, 4, 5})
	assert.Equal(t, 120, prod[0])
	assert.Equal(t, 60, prod[1])
	assert.Equal(t, 40, prod[2])
	assert.Equal(t, 30, prod[3])
	assert.Equal(t, 24, prod[4])
}

func TestExProd2(t *testing.T) {
	prod := exprod.Product([]int{3, 2, 1})
	assert.Equal(t, 2, prod[0])
	assert.Equal(t, 3, prod[1])
	assert.Equal(t, 6, prod[2])
}
