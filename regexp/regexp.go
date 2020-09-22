// Package regexp contains Daily Coding Problem #25
// This problem was asked by Facebook.
// Implement regular expression matching with the following special characters:
// . (period) which matches any single character
// * (asterisk) which matches zero or more of the preceding element
// That is, implement a function that takes in a string and a valid
// regular expression and returns whether or not the string matches
// the regular expression.
// For example, given the regular expression "ra." and the string "ray",
// your function should return true. The same regular expression on
// the string "raymond" should return false.
// Given the regular expression ".*at" and the string "chat",
// your function should return true. The same regular expression on
// the string "chats" should return false.
package regexp

// Check checks if str satisfies regular expression regexp
// containing . and *
func Check(str, regexp string) bool {
	if str == "" && (regexp == "" || regexp == "*") {
		return true
	}
	if str == "" || regexp == "" {
		return false
	}

	if regexp[0] == '.' {
		return Check(str[1:], regexp[1:])
	}

	if regexp[0] == '*' {
		return Check(str[1:], regexp[1:]) || Check(str[1:], regexp)
	}

	if regexp[0] != str[0] {
		return false
	}

	return Check(str[1:], regexp[1:])
}
