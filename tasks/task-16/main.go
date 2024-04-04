package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{8, 9, 10, 11, 12, 13, 14, 15, 16, 1, 2, 3, 4, 5, 6, 7}

	slices.Sort(nums)

	fmt.Println(nums)
}
