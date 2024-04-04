package main

import (
	"fmt"
	"sync"
)

const workersCount = 3

func main() {
	concurrencySquare([]int{2, 4, 6, 8, 10})
	workerSquare([]int{2, 4, 6, 8, 10})
}

// вариант с использованием workerPool
func workerSquare(nums []int) {
	var wg sync.WaitGroup
	wg.Add(workersCount)
	inputCh := make(chan int, workersCount)

	for i := 0; i < workersCount; i++ {
		go func(outCh <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for val := range outCh {
				fmt.Println(val * val)
			}
		}(inputCh, &wg)
	}

	for _, val := range nums {
		inputCh <- val
	}

	close(inputCh)
	wg.Wait()
}

// горутина на каждый эллемент
func concurrencySquare(nums []int) {
	var wg sync.WaitGroup

	for _, v := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(v)
	}

	wg.Wait()
}
