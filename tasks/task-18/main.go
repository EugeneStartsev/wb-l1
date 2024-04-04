package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counterByMutex struct {
	mutex sync.Mutex
	i     int
}

const workers = 10

func main() {
	valByMutex := counterByMutex{i: 0}

	wg := &sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go workerByMutex(wg, &valByMutex)
	}

	wg.Wait()
	fmt.Println(valByMutex.i)

	var valByAtomic int64 = 0

	wg = &sync.WaitGroup{}
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go workerByAtomic(wg, &valByAtomic)
	}

	wg.Wait()
	fmt.Println(valByAtomic)
}

func workerByAtomic(wg *sync.WaitGroup, val *int64) {
	defer wg.Done()

	for i := 0; i < 50; i++ {
		atomic.AddInt64(val, 1)
	}
}

func workerByMutex(wg *sync.WaitGroup, c *counterByMutex) {
	defer wg.Done()

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := 0; i < 50; i++ {
		c.i++
	}
}
