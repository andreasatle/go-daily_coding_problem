// Package encode contains Daily Coding Problem #7
//
// This problem was asked by Facebook.
//
// Given the mapping a = 1, b = 2, ... z = 26, and an encoded message,
// count the number of ways it can be decoded.
// For example, the message '111' would give 3,
// since it could be decoded as 'aaa', 'ka', and 'ak'.
// You can assume that the messages are decodable.
// For example, '001' is not allowed.
//
// Note: Originally, the decoded messages was also outputted.
// At a price of a more complicated code.
//
// Author: Andreas Atle, atle.andreas@gmail.com
package encode

import (
	"errors"
	"strconv"
	"strings"
)

const (
	// charOffset defines the mapping {'a':1, 'b':2, ...}
	charOffset = int('a') - 1
)

// Encode encode a string of lower letters to a string of numbers
func Encode(msg string) (string, error) {
	slice := []string{}
	for _, ch := range msg {
		if ch < 'a' || ch > 'z' {
			return "", errors.New("found a non-lower case character")
		}
		slice = append(slice, strconv.Itoa(int(ch)-charOffset))
	}
	return strings.Join(slice, ""), nil
}

// DecodeCount counts the possible number of decoded messages.
func DecodeCount(enc string) int {
	// Decoding finished, print to screen (for now)
	if enc == "" {
		return 1
	}

	// Invalid case, enc[0] == '0' is not allowed
	if rune(enc[0]) == '0' {
		return 0
	}

	enc1 := enc[1:]
	counter := 0

	// Decode entry 1-9
	counter += DecodeCount(enc1)

	// Decode 10-26
	if len(enc) >= 2 && (rune(enc[0]) == '1' || (rune(enc[0]) == '2' && rune(enc[1]) >= '0' && rune(enc[1]) <= '6')) {
		enc1 := enc[2:]
		counter += DecodeCount(enc1)
	}

	return counter
}
