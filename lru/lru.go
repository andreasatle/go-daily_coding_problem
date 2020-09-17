// Package lru contains Daily Coding Problem #52
// This problem was asked by Google.
// Implement an LRU (Least Recently Used) Cache.
// It should be able to be initialized with a Cache size n,
// and contain the following methods:
// set(key, value): sets key to value.
// - If there are already n items in the Cache and we are adding a new item,
// - then it should also remove the least recently used item.
// get(key): gets the value at key. If no such key exists, return null.
// - Each operation should run in O(1) time.
package lru

import "fmt"

// Key is the type of the key
type Key string

// Value is the type of the value
type Value string

// Cache contains at most size (key,value)-pairs with keys also stored in Queue
type Cache struct {
	curr  int
	size  int
	Queue []Key
	Cache map[Key]*Value
}

// NewCache creates a new Cache with given size
func NewCache(size int) *Cache {
	lru := &Cache{curr: 0, size: size}
	lru.Queue = make([]Key, size)
	lru.Cache = map[Key]*Value{}
	return lru
}

// Get the value given a key.
func (lru *Cache) Get(key Key) *Value {
	val, ok := lru.Cache[key]
	if !ok {
		return nil
	}
	return val
}

// Set a (key,value)-pair.
// If maximum size is reached, then delete the current entry in the Queue from the Cache.
// If the Cache updates the value for an existing key, the Queue is not affected,
// and the new Cache-pair will appear older than it is.
func (lru *Cache) Set(key Key, value Value) {
	lenBefore := len(lru.Cache)
	lru.Cache[key] = &value
	lenAfter := len(lru.Cache)
	if lenAfter > lenBefore {
		delete(lru.Cache, lru.Queue[lru.curr])
		lru.Queue[lru.curr] = key
		lru.curr = (lru.curr + 1) % lru.size
	}
}

// Print outputs Cache to stdout
func (lru *Cache) Print() {
	fmt.Printf("Cache-size: %d\n", lru.size)
	fmt.Printf("Current: %d\n", lru.curr)
	fmt.Printf("Queue: %v\n", lru.Queue)
	fmt.Printf("Cache: ")
	str := ""
	for k, v := range lru.Cache {
		fmt.Print(str + "(" + string(k) + "," + string(*v) + ")")
		str = ", "
	}
	fmt.Println()

}
