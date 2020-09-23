package cons_test

import (
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/cons"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	pair := cons.Cons(3, 4)
	assert.Equal(t, 3, cons.Car(pair))
	assert.Equal(t, 4, cons.Cdr(pair))
}
