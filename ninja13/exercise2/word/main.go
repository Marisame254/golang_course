// Package word provides custom functions for working with strings and words.
package word

import "strings"

// UseCount return the number of tines each word used in a string
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// count return the numbers in the string
func Count(s string) int {
	// write the code for this func
	xs := strings.Fields(s)
	return len(xs)
}
