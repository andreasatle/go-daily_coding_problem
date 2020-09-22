// Package rnd contains Daily Coding Problem #90
// This question was asked by Google.
// Given an integer n and a list of integers l,
// write a function that randomly generates a number from
// 0 to n-1 that isn't in l (uniform).
package rnd

import (
	"errors"
	"math/rand"
	"time"
)

// initRnd returns a slice and length of entries not in list.
func initRnd(n int, list []int) ([]int, int) {
	m := map[int]struct{}{}
	for _, l := range list {
		m[l] = struct{}{}
	}
	mapping := []int{}
	for i := 0; i < n; i++ {
		if _, ok := m[i]; !ok {
			mapping = append(mapping, i)
		}
	}
	return mapping, len(mapping)
}

// Create returns a random generator function.
func Create(n int, list []int) (func() int, error) {

	// Set random seed by current time.
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	mapping, length := initRnd(n, list)

	// Create the random generator function to be returned.
	rand := func() int {
		return mapping[r.Intn(length)]
	}

	if len(mapping) == 0 {
		return nil, errors.New("no values to choose from")
	}

	return rand, nil
}
