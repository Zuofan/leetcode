package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func isAnagram(first, second string) bool {
	//(2)
	f := strings.Split(first, "")
	sort.Strings(f)
	firstNew := strings.Join(f, "")

	s := strings.Split(second, "")
	sort.Strings(s)
	secondNew := strings.Join(s, "")

	return strings.Compare(firstNew, secondNew) == 0
}


type sortBytes []byte

func (s sortBytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortBytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortBytes) Len() int {
	return len(s)
}

func isAnagramV2(first, second string) bool {
	f := []byte(first)
	s := []byte(second)
	
	sort.Sort(sortBytes(f))
	sort.Sort(sortBytes(s))

	return bytes.Compare(f, s) == 0
}

func main() {
	var first, second string

	fmt.Printf("Enter two strings and I'll tell you if they are anagrams:\n ")
	fmt.Printf("Enter the first string: ")
	fmt.Scanf("%s", &first)
	fmt.Printf("Enter the second string: ")
	fmt.Scanf("%s", &second)

	//(1)
	if len(first) != len(second) {
		fmt.Printf("%s and %s are not anagrams.\n", strconv.Quote(first), strconv.Quote(second))
		return
	}

	//(2)
	if isAnagram(first, second) {
		fmt.Printf("%s and %s are anagrams.\n", strconv.Quote(first), strconv.Quote(second))
	} else {
		fmt.Printf("%s and %s are not anagrams.\n", strconv.Quote(first), strconv.Quote(second))
	}
}