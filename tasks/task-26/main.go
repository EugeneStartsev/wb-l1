package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcd"
	fmt.Println(checkDuplicate(s))

	s = "abCdefAaf"
	fmt.Println(checkDuplicate(s))

	s = "aabcd"
	fmt.Println(checkDuplicate(s))
}

func checkDuplicate(s string) bool {
	s = strings.ToLower(s)

	m := make(map[rune]int)

	for _, val := range s {
		m[val]++

		if m[val] > 1 {
			return false
		}
	}

	return true
}
