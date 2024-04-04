package main

import "fmt"

func main() {
	something := []string{"cat", "cat", "cat", "dog", "tree"}
	m := make(map[string]struct{})

	var set []string

	for _, v := range something {
		if _, ok := m[v]; !ok {
			set = append(set, v)
			m[v] = struct{}{}
		}
	}

	fmt.Println(set)
}
