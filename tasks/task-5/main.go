package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var delay time.Duration
	_, err := fmt.Scan(&delay)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)

	ctx, cancel := context.WithTimeout(context.Background(), delay*time.Second)
	defer func() {
		fmt.Println("ctx done")
		cancel()
	}()

	go func(ch <-chan int, wg *sync.WaitGroup) {
		for {
			select {
			case val := <-ch:
				fmt.Println(val)
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}(ch, &wg)

	go func(ch chan<- int) {
		for {
			var data int
			select {
			case <-time.After(delay):
				_, err = fmt.Scan(&data)
				if err != nil {
					log.Fatal(err)
				}

				ch <- data
			case <-ctx.Done():
				return
			}
		}
	}(ch)

	wg.Wait()
}
