//
//	Package app provide concurrency examples.
//
package app

import (
	"time"
	"fmt"
)

// start worker then wait from response from worker,
// otherwise time out in 1 second.
func WaitResponse() {
	fmt.Println("WaitResponse")

	timeout := make(chan bool, 1)
	ch := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	go Worker(ch)

	select {
	case x := <-ch:
		fmt.Println(x)
	case <-timeout:
		fmt.Println("time out")
		return
	}
}

func WaitResponse2() {

	fmt.Println("WaitResponse2")
	ch := make(chan int)

	go Worker(ch)

	select {
	case x := <-ch:
		fmt.Println(x)
	case <-time.After(1 * time.Second):
		fmt.Println("time out")
		return
	}

}

func Worker(ch chan int) {
	ch <- 1
	time.Sleep(2 * time.Second)
}
