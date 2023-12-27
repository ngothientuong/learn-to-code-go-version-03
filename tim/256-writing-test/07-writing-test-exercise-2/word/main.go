// Package word contains function to count the word of a long string
package word

import (
	"strings"
)

// Function UseCount counts keep tracks of how many times a word in a string is repeated
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	// Split the string into slice of string with white space delimiter
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Function Count return the word counts of the string
func Count(s string) int {
	// write the code for this func
	// Split the string into slice of string with white space delimiter
	xs := strings.Fields(s)
	return len(xs)
}
