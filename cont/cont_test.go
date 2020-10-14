package cont_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/cont"
	"github.com/stretchr/testify/assert"
)

func TestCont(t *testing.T) {
	var terms []int

	// Working case
	terms = cont.FindTerms([]int{2, 3, 6, 9, -3, 4, 1}, 10)
	assert.Equal(t, 9, terms[0])
	assert.Equal(t, -3, terms[1])
	assert.Equal(t, 4, terms[2])

	// Empty case
	terms = cont.FindTerms([]int{}, 10)
	assert.Equal(t, 0, len(terms))

	// Not working case
	terms = cont.FindTerms([]int{}, 10)
	assert.Equal(t, 0, len(terms))
}
