// Package rgb contains Daily Coding Problem #35
// This problem was asked by Google.
// Given an array of strictly the characters 'R', 'G', and 'B',
// segregate the values of the array so that all the Rs come first,
// the Gs come second, and the Bs come last. You can only swap
// elements of the array.
// Do this in linear time and in-place.
// For example, given the array ['G', 'B', 'R', 'R', 'B', 'R', 'G'],
// it should become ['R', 'R', 'R', 'G', 'G', 'B', 'B'].
package rgb

type empty struct{}

// Sort sorts a slice of 'R', 'G', 'B' runes
func Sort(rgb []rune) {
	nextR := 0
	nextB := len(rgb) - 1
	for i := nextR; i <= nextB; {
		if rgb[i] == 'R' {
			rgb[nextR], rgb[i] = rgb[i], rgb[nextR]
			nextR++
		} else if rgb[i] == 'B' {
			rgb[nextB], rgb[i] = rgb[i], rgb[nextB]
			nextB--
		} else {
			i++
		}
	}
}
