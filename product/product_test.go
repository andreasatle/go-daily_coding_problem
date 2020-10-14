package product_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/product"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	assert.Equal(t, 500, product.LargestProductOfThree([]int{-10, -10, 5, 2}), "Wrong largest product of three.")
	assert.Equal(t, 500, product.LargestProductOfThree([]int{10, 10, 5, 2}), "Wrong largest product of three.")
	assert.Equal(t, 0, product.LargestProductOfThree([]int{-1, -2}), "Wrong largest product of three.")
	assert.Equal(t, -6, product.LargestProductOfThree([]int{-1, -2, -4, -3}), "Wrong largest product of three.")

}
