// Package singleton contains Daily Coding Problem #120
//
// This problem was asked by Microsoft.
//
// Implement the singleton pattern with a twist. First, instead of storing
// one instance, store two instances. And in every even call of getInstance(),
// return the first instance and in every odd call of getInstance(),
// return the second instance.
//
// Note: The problem is very vague, it is unclear what the instance is
// an instance of...
//
// Author: Andreas Atle, atle.andreas@gmail.com
package singleton

type Singleton struct {
	SomeFunData string
}

// New returns a function that returns a pointer to one of two singletons.
func New() func() *Singleton {
	flag := true
	oddData := &Singleton{SomeFunData: "This is odd!"}
	evenData := &Singleton{SomeFunData: "Time to get even!"}

	getInstance := func() *Singleton {
		defer func() { flag = !flag }()
		if flag {
			return oddData
		}
		return evenData
	}

	return getInstance
}
