package lru_test

import (
	"testing"

	"github.com/andreasatle/Assorted/Daily/lru"
	"github.com/stretchr/testify/assert"
)

func TestLruCache(t *testing.T) {

	lruCache := lru.NewCache(3)

	lruCache.Set("foo", "bar")
	lruCache.Set("bar", "foo")
	lruCache.Set("foo", "bar1")

	assert.Equal(t, "bar1", string(*lruCache.Get("foo")), "Duplicates not handled correctly.")
	assert.Equal(t, 2, len(lruCache.Cache), "Cache of wrong size.")

	lruCache.Set("John", "Doe")
	lruCache.Set("Steve", "Doe")
	lruCache.Set("png", "jpg")

	assert.Equal(t, 3, len(lruCache.Queue), "Wrong size of queue.")
	assert.Equal(t, 3, len(lruCache.Cache), "Wrong size of cache.")
	assert.Equal(t, "jpg", string(*lruCache.Get("png")), "Wrong (key,value)-pair.")
}
