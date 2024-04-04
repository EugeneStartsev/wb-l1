package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	channelSum([]int{2, 4, 6, 8, 10})
	atomicSum([]int{2, 4, 6, 8, 10})
}

// через атомики
func atomicSum(nums []int) {
	var wg sync.WaitGroup
	var res int64

	for _, v := range nums {
		wg.Add(1)
		go func(v int64) {
			defer wg.Done()
			atomic.AddInt64(&res, v*v)
		}(int64(v))
	}

	wg.Wait()

	fmt.Println(res)
}

// через каналы
func channelSum(nums []int) {
	var wg sync.WaitGroup
	ch := make(chan int, len(nums))
	var res int

	for _, v := range nums {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()

			sum := 0
			sum += v * v

			ch <- sum
		}(v)
	}

	for i := 0; i < len(nums); i++ {
		res += <-ch
	}

	fmt.Println(res)

	wg.Wait()
}
