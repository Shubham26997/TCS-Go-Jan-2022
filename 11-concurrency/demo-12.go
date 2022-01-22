/* Using range construct on channels */
package main

import (
	"fmt"
	"time"
)

func main() {
	noCh := make(chan int)
	go genFibonacci(noCh)
	for no := range noCh {
		fmt.Println(no)
	}
}

func genFibonacci(resultCh chan int) {
	x, y := 0, 1
	doneCh := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		doneCh <- "stop"
	}()

	for {
		select {
		case <-doneCh:
			close(resultCh)
			return
		case resultCh <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		}
	}
}
