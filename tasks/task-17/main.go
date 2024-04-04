package main

import "fmt"

func binarySearch(nums []int, i int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == i {
			return mid
		}
		if nums[mid] > i {
			r = mid - 1
		}
		if nums[mid] < i {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	nums := []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearch(nums, 7))
}
