// Package substr contains Daily Coding Problem #103
//
// This problem was asked by Square.
//
// Given a string and a set of characters, return the shortest substring
// containing all the characters in the set.
//
// For example, given the string
//  "figehaeci"
// and the set of characters
//  {a, e, i},
// you should return
//  "aeci".
//
// If there is no substring containing all the characters in the set, return null.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package substr

// Empty is a building block of a set
type Empty struct{}

// ByteSet is a set of bytes
type ByteSet map[byte]Empty

// NewByteSet returns a new instance of ByteSet (no pointer)
func NewByteSet(str string) ByteSet {
	byteSet := ByteSet{}
	for i := 0; i < len(str); i++ {
		(byteSet)[str[i]] = Empty{}
	}
	return byteSet
}

// SubSet returns the smallest substring that contains all bytes in the ByteSet
func (bytes ByteSet) SubSet(str string) (string, bool) {
	first, last, ok := bytes.subSet(str, 0, len(str)-1)
	if ok {
		return str[first : last+1], ok
	}
	return "", ok
}

func (bytes ByteSet) subSet(str string, first, last int) (int, int, bool) {
	if !bytes.contains(str, first, last) {
		return -1, -1, false
	}

	first1, last1, ok1 := bytes.subSet(str, first, last-1)
	first2, last2, ok2 := bytes.subSet(str, first+1, last)

	if !ok2 && ok1 {
		return first1, last1, ok1
	}

	if !ok1 && ok2 {
		return first2, last2, ok2
	}

	// Both ok1 and ok2 are true
	if first1 != -1 && last1-first1 <= last2-first2 {
		return first1, last1, ok1
	}
	if first2 != -1 && last2-first2 <= last1-first1 {
		return first2, last2, ok2
	}

	return first, last, true
}

func (bytes ByteSet) contains(str string, first, last int) bool {
	bytes2 := ByteSet{}
	for i := first; i <= last; i++ {
		if _, ok := bytes[str[i]]; ok {
			bytes2[str[i]] = Empty{}
		}
	}
	return len(bytes) == len(bytes2)
}
