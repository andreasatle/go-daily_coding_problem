// Package recent contains Daily Coding Problem #97
//
// This problem was asked by Stripe.
//
// Write a map implementation with a get function that lets you retrieve
// the value of a key at a particular time.
//
// It should contain the following methods:
//  set(key, value, time): sets key to value for t = time.
//  get(key, time): gets the key at t = time.
// The map should work like this. If we set a key at a particular time,
// it will maintain that value forever or until it gets set at a later
// time. In other words, when we get a key at a time, it should return
// the value that was set for that key set at the most recent time.
//
// Consider the following examples:
//  d.set(1, 1, 0) # set key 1 to value 1 at time 0
//  d.set(1, 2, 2) # set key 1 to value 2 at time 2
//  d.get(1, 1)    # get key 1 at time 1 should be 1
//  d.get(1, 3)    # get key 1 at time 3 should be 2
//  d.set(1, 1, 5) # set key 1 to value 1 at time 5
//  d.get(1, 0)    # get key 1 at time 0 should be null
//  d.get(1, 10)   # get key 1 at time 10 should be 1
//  d.set(1, 1, 0) # set key 1 to value 1 at time 0
//  d.set(1, 2, 0) # set key 1 to value 2 at time 0
//  d.get(1, 0)    # get key 1 at time 0 should be 2
//
// Author: Andreas Atle, atle.andreas@gmail.com
package recent

import (
	"errors"
	"sort"
)

type Tuple struct {
	Value int
	Time  int
}

type Map map[int][]Tuple

func NewMap() Map {
	return Map{}
}

func (m Map) Set(key, value, time int) {

	// First entry of m[key]
	if _, ok := m[key]; !ok {
		m[key] = []Tuple{Tuple{Value: value, Time: time}}
		return
	}

	// m[key].Time == time already exist, overwrite m[key].Value
	cmp := func(i int) bool {
		return m[key][i].Time >= time
	}
	ind := sort.Search(len(m[key]), cmp)
	if ind < len(m[key]) && m[key][ind].Time == time {
		m[key][ind].Value = value
		return
	}

	// Insert last and reverse bubble sort
	m[key] = append(m[key], Tuple{Value: value, Time: time})
	for i := len(m[key]) - 1; i >= 1; i-- {
		if m[key][i].Time > m[key][i-1].Time {
			break
		}
		m[key][i-1], m[key][i] = m[key][i], m[key][i-1]
	}
}

func (m Map) Get(key, time int) (int, error) {
	if _, ok := m[key]; !ok {
		return 0, errors.New("Error - key not found")
	}

	cmp := func(i int) bool {
		return m[key][i].Time >= time
	}
	ind := sort.Search(len(m[key]), cmp)
	if ind < len(m[key]) && m[key][ind].Time == time {
		return m[key][ind].Value, nil
	}

	if ind == 0 {
		return 0, errors.New("Error - retrieving data before timestamp")
	}

	return m[key][ind-1].Value, nil
}
