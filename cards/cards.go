// Package cards contains Daily Coding Problem #51
//
// This problem was asked by Facebook.
//
// Given a function that generates perfectly random numbers between 1 and k
// (inclusive), where k is an input, write a function that shuffles a deck
// of cards represented as an array using only swaps.
//
// It should run in O(N) time.
//
// Hint: Make sure each one of the 52! permutations of the deck is
// equally likely.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package cards

import (
	"math/rand"
	"time"
)

// Shuffle returns a shuffled deck of cards
func Shuffle(n int) []int {
	deck := make([]int, n)
	for i := range deck {
		deck[i] = i // start with index 0
	}

	rand := createRand(n)
	for i := 0; i < n; i++ {
		j := rand()
		deck[i], deck[j] = deck[j], deck[i]
	}
	return deck

}

// rand computes a random number [0, n-1]. (C- rather than F-notation)
func createRand(n int) func() int {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	n64 := float64(n)

	rand := func() int {
		return int(rnd.Float64() * n64)
	}

	return rand
}
