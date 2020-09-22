package product_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/product"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	expected := 500
	actual := product.LargestProductOfThree([]int{-10, -10, 5, 2})
	assert.Equal(t, expected, actual, "Wrong largest product of three.")
}
