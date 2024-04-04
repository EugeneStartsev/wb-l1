package main

import (
	"fmt"
	"time"
)

func sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	t := time.Now()
	sleep(time.Second * 10)
	fmt.Println(time.Since(t))
}
