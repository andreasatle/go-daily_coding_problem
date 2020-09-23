// Package cons contains Daily Coding Problem #5
// This problem was asked by Jane Street.
// cons(a, b) constructs a pair, and car(pair) and cdr(pair) returns the first and last element of that pair. For example, car(cons(3, 4)) returns 3, and cdr(cons(3, 4)) returns 4.
// Given this implementation of cons:
// def cons(a, b):
//     def pair(f):
//         return f(a, b)
//     return pair
// Implement car and cdr.
// Note: This is functional madness! What is the point?!
package cons

// Func2 is a function Z x Z => Z
type Func2 func(int, int) int

// Pair is a function (Z x Z => Z) => Z
type Pair func(Func2) int

// Cons defines a tuple in a very weird way
func Cons(a, b int) Pair {
	pair := func(f Func2) int {
		return f(a, b)
	}
	return pair
}

// Car recover the first entry in the tuple
func Car(pair Pair) int {
	return pair(func(x, y int) int { return x })
}

// Cdr recover the second entry in the tuple
func Cdr(pair Pair) int {
	return pair(func(x, y int) int { return y })
}
