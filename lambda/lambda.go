// Package lambda contains Daily Coding Problem #91
//
// This problem was asked by Dropbox.
//
// What does the below code snippet print out? How can we fix the anonymous functions to behave as we'd expect?
//  functions = []
//  for i in range(10):
//      functions.append(lambda : i)
//  for f in functions:
//      print(f())
//
// Note: Lexical scoping gives that f() returns
// the value of i in the current scope.
// We need to wrap the functions to get the
// expected behaviour.
//	get, acc, set, lessThan := lambda.CreateNext(0)
//	for set(1); lessThan(4); acc() {
//		fmt.Printf("%d ", get())
//	}
// prints "1 2 3 "
//
// Author: Andreas Atle, atle.andreas@gmail.com
package lambda

// CreateNext creates a "Lexical scope" to create an iterator
func CreateNext(i int) (func() int, func(), func(int), func(int) bool) {

	// setValue sets the first number to be returned by nextValue
	setValue := func(j int) {
		i = j
	}

	// accValue updates the local variable to the next value
	accValue := func() {
		i++
	}

	// lessThan compares the local variable with input arg
	lessThan := func(n int) bool {
		return i < n
	}

	// nextValue returns and updates (after) the local value
	getValue := func() int {
		return i
	}
	return getValue, accValue, setValue, lessThan
}
