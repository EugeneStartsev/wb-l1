package main

import "fmt"

func main() {
	fmt.Println(deleteFromSlice([]int{1, 2, 3, 4, 5, 6}, 2))
}

func deleteFromSlice(nums []int, i int) []int {
	copy(nums[i:], nums[i+1:])
	return nums[:len(nums)-1]
}
