// Package sumprimes contains Daily Coding Problem #101
//
// This problem was asked by Alibaba.
//
//  Given an even number (greater than 2),
//  return two prime numbers whose sum will be equal to the given number.
// A solution will always exist. See Goldbach’s conjecture.
//
// Example:
//  Input: 4
//  Output: 2 + 2 = 4
// If there are more than one solution possible,
// return the lexicographically smaller solution.
//
// If [a, b] is one solution with a <= b,
// and [c, d] is another solution with c <= d, then
//  [a, b] < [c, d]
//  If a < c OR a==c AND b < d.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package sumprimes

import "errors"

// PrimeSum returns to primes that sum up to the argument.
// When there are multiple solutions, the one with the smallest
// prime number is chosen.
func PrimeSum(n int) (int, int, error) {
	if n%2 == 1 {
		return 0, 0, errors.New("odd input not allowed")
	}
	primes := getPrimesUpTo(n)
	for first, last := 0, len(primes)-1; first <= last; {
		if p1, p2 := primes[first], primes[last]; p1+p2 == n {
			return p1, p2, nil
		} else if p1+p2 < n {
			first++
		} else {
			last--
		}
	}
	return 0, 0, errors.New("no primes found")
}

func getPrimesUpTo(n int) []int {
	if n < 2 {
		return []int{}
	}

	primes := []int{}
	for i := 2; i <= n; i++ {
		if numDividedByPrime(i, primes) {
			continue
		}
		primes = append(primes, i)
	}
	return primes
}

func numDividedByPrime(n int, primes []int) bool {
	for _, prime := range primes {
		if n%prime == 0 {
			return true
		}
	}
	return false
}
