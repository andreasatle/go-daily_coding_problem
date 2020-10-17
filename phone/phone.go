// Package phone contains Daily Coding Problem #81
//
// This problem was asked by Yelp.
//
// Given a mapping of digits to letters (as in a phone number), and a digit
// string, return all possible letters the number could represent. You can
// assume each valid number in the mapping is a single digit.
//
// For example if {“2”: [“a”, “b”, “c”], 3: [“d”, “e”, “f”], …} then “23”
// should return [“ad”, “ae”, “af”, “bd”, “be”, “bf”, “cd”, “ce”, “cf"].
//
// Author: Andreas Atle, atle.andreas@gmail.com
package phone

// Runes is a rune-slice
type Runes []rune

// PhoneEncoder maps a rune to a slice of runes
type PhoneEncoder map[rune]Runes

// Encode returns all possible letter versions of a phone number
func Encode(number string) []string {
	pe := newPhoneEncoder(number)

	out := []string{}

	if len(number) == 0 {
		return out
	}

	for _, chr := range pe[rune(number[0])] {
		out = append(out, string(chr))
	}

	for i := 1; i < len(number); i++ {
		r := rune(number[i])
		out2 := out
		out = []string{}
		for _, str := range out2 {
			for _, chr := range pe[r] {
				out = append(out, str+string(chr))
			}
		}

	}
	return out
}

func newPhoneEncoder(number string) PhoneEncoder {
	phoneEncoder := PhoneEncoder{}
	phoneEncoder['0'] = Runes{'0'}
	phoneEncoder['1'] = Runes{'1'}
	phoneEncoder['2'] = Runes{'a', 'b', 'c'}
	phoneEncoder['3'] = Runes{'d', 'e', 'f'}
	phoneEncoder['4'] = Runes{'g', 'h', 'i'}
	phoneEncoder['5'] = Runes{'j', 'k', 'l'}
	phoneEncoder['6'] = Runes{'m', 'n', 'o'}
	phoneEncoder['7'] = Runes{'p', 'q', 'r', 's'}
	phoneEncoder['8'] = Runes{'t', 'u', 'v'}
	phoneEncoder['9'] = Runes{'w', 'x', 'y', 'z'}

	for _, ch := range number {
		r := rune(ch)
		if _, ok := phoneEncoder[r]; !ok {
			phoneEncoder[r] = Runes{r}
		}
	}
	return phoneEncoder
}
