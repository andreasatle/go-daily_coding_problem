package xor_test

import (
	"fmt"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/xor"
	"github.com/stretchr/testify/assert"
)

func TestXor(t *testing.T) {

	array := []xor.Value{4711, 3500, 8913, 1174}
	lst := &xor.DLinkedList{}
	for _, val := range array {
		// The routine returns the inserted node, but is not used
		// Unclear how garbage collection is dealing with this.
		// This has to be thoroughly investigated;)
		lst.Insert(val)
	}
	fmt.Println(lst.Slice())

	// Get inbounds indices
	for i := 0; i < 4; i++ {
		n, err := lst.Get(i)
		if err != nil {
			assert.Equal(t, true, false, "Out of bounds.")
			break
		}
		assert.Equal(t, array[i], n.Val)
		fmt.Println(i, n.Val)
	}
	// Get an out of bounds
	_, err := lst.Get(5)
	assert.Equal(t, true, err != nil)
}
