package app

import (
	"testing"
	"fmt"
)

func Test_sum(t *testing.T) {
	request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
	q := make(chan *Request,1)
	q <- request
	go handle2(q)
	fmt.Printf("answer: %d\n", <-request.resultChan)
}

func handle2(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}