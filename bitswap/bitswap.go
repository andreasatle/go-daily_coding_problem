// Package bitswap contains Daily Coding Problem #109
//
// This problem was asked by Cisco.
//
// Given an unsigned 8-bit integer, swap its even and odd bits.
// The 1st and 2nd bit should be swapped, the 3rd and 4th bit
// should be swapped, and so on.
//
// For example, 10101010 should be 01010101. 11100010 should be 11010001.
//
// Bonus: Can you do this in one line?
//
// Author: Andreas Atle, atle.andreas@gmail.com
package bitswap

const (
	evenMask uint8 = 0xAA
	oddMask  uint8 = 0x55
)

// Apply swaps even and odd digits in an unsigned 8-bit int
func Apply(num uint8) uint8 {
	return (evenMask&num)>>1 + (oddMask&num)<<1
}
