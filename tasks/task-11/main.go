package main

import "fmt"

func main() {
	set := []int{1, 2, 3, 4, 5, 6}
	set2 := []int{3, 4, 5, 8, 9}

	fmt.Println(crossSet(set, set2))
}

func crossSet(set, set2 []int) []int {
	m := make(map[int]int)

	res := make([]int, 0, len(set))

	for _, val := range set {
		m[val] += 1
	}

	for _, val := range set2 {
		m[val] += 1
	}

	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}

	return res
}
