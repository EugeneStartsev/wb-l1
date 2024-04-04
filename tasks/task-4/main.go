package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func startWorker(ctx context.Context, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Printf("worker read: %s\n", val)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	var workersCount int

	_, err := fmt.Scan(&workersCount)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(workersCount)

	ch := make(chan string, workersCount)
	defer close(ch)

	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < workersCount; i++ {
		go startWorker(ctx, ch, &wg)
	}

	delay := time.Second * 5

	go func() {
		for {
			var data string

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
	}()

	signch := make(chan os.Signal)
	signal.Notify(signch, os.Interrupt)
	<-signch
	cancel()
	wg.Wait()
}
