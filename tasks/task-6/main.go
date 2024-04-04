package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func channelClose(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	_, ok := <-ch
	if !ok {
		fmt.Println("goroutine finished")
	}
}

func rangeChannelCLose(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("goroutine finished by range")
}

func channelCloseByStopChannel(ch chan int, stop chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-stop:
			fmt.Println("goroutine finished by stop channel")
			return
		}
	}
}

func channelCloseByContext(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("goroutine finished by context")
			return
		}
	}
}

func channelCloseByTimeoutContext(ctx context.Context, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("goroutine finished by timeout")
			return
		}
	}
}

/*func channelCloseByPanic(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	close(ch)
	fmt.Println("goroutine finished by panic")
	ch <- 1
}*/

func main() {
	ch := make(chan int)

	wg := &sync.WaitGroup{}
	wg.Add(5)

	close(ch)
	go channelClose(ch, wg)

	ch = make(chan int, 2)
	ch <- 3
	ch <- 4
	close(ch)
	go rangeChannelCLose(ch, wg)

	ch = make(chan int)
	stop := make(chan struct{})
	go channelCloseByStopChannel(ch, stop, wg)
	stop <- struct{}{}

	ctx, cancel := context.WithCancel(context.Background())
	go channelCloseByContext(ctx, ch, wg)
	cancel()

	ctxTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	go channelCloseByTimeoutContext(ctxTimeout, ch, wg)

	//go channelCloseByPanic(ch, wg)

	wg.Wait()
}
