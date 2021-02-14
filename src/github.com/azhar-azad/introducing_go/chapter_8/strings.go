package main

import (
	"fmt"
	"strings"
)

func stringsDemo() {

	// search for a smaller string in a bigger string.
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es"))
	// => true

	// count the number of times a smaller string occurs in a bigger string.
	// Count(s, sep string) int
	fmt.Println(strings.Count("test", "t"))
	// => 2

	// determine if a bigger string starts with a smaller string.
	// HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))
	// => true

	// determine if a bigger string ends with a smaller string.
	// HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("test", "st"))
	// => true

	// find the position of a smaller string in a bigger string.
	// Index(s, sep string) int
	fmt.Println(strings.Index("test", "e"))
	// => 1

	// take a list of strings and join them together in a single string
	// separated by another string (e.g., a comma).
	// Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	// => "a-b"

	// repeat a string.
	// Repeat(s string, count int) string
	fmt.Println(strings.Repeat("a", 5))
	// => aaaaa

	// replace a smaller string in a bigger string with some other string.
	// Replace also takes a number indicating how many times to do the
	// replacement (pass -1 to do it as many times as possible)
	// Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	// => "bbaa"

	// split a string into a list of strings by a separating string.
	// Split is the reverse of Join
	// Split(s, sep string) []string
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	// => [a b c d e]

	// convert a string to all lowercase letters.
	// ToLower(s string) string
	fmt.Println(strings.ToLower("TEST"))
	// => "test"

	// convert a string to all uppercase letters.
	// ToUpper(s string) string
	fmt.Println(strings.ToUpper("test"))
	// => "TEST"

	// convert a string to a slice of bytes (and vice versa).
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println("arr = ", arr)
	fmt.Println("str = ", str)
	// 	arr =  [116 101 115 116]
	//  str =  test
}
