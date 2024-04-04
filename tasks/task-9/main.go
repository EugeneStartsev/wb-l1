package main

import (
	"fmt"
	"sync"
)

func main() {
	genFirst()

	ch := generateNumbers([]int{1, 2, 3, 4, 5, 6, 7})
	outputch := calculateSqr(ch)

	for val := range outputch {
		fmt.Println(val)
	}
}

func genFirst() {
	inputch := make(chan int)
	outputch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	wg2 := &sync.WaitGroup{}
	wg2.Add(1)

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	go func() {
		defer wg.Done()

		for val := range inputch {
			outputch <- val * val
		}
	}()

	go func() {
		defer wg2.Done()
		for val := range outputch {
			fmt.Println(val)
		}
	}()

	for _, val := range nums {
		inputch <- val
	}

	close(inputch)
	wg.Wait()
	close(outputch)
	wg2.Wait()
}

func generateNumbers(nums []int) chan int {
	outputch := make(chan int)

	go func() {
		for _, val := range nums {
			outputch <- val
		}

		close(outputch)
	}()

	return outputch
}

func calculateSqr(input chan int) chan int {
	outputch := make(chan int)

	go func() {
		for val := range input {
			outputch <- val * val
		}

		close(outputch)
	}()

	return outputch
}
