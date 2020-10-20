package cards_test

import (
	"sort"
	"testing"

	"github.com/andreasatle/go-daily_coding_problem/cards"
	"github.com/stretchr/testify/assert"
)

func TestCards(t *testing.T) {
	helpFun(t, 40)
	helpFun(t, 52)
}

func helpFun(t *testing.T, n int) {

	deck := cards.Shuffle(n)
	deckBkp := make([]int, len(deck))
	copy(deckBkp, deck)
	sort.Ints(deck)

	nDifferent := 0
	for i := 0; i < n; i++ {
		assert.Equal(t, i, deck[i])
		if i != deckBkp[i] {
			nDifferent++
		}
	}

	// Assume that atleast 90% of all cards are not sorted
	// This doesn't work when there are too few cards (<20 or so)
	assert.GreaterOrEqual(t, nDifferent*10, n*9)
}
