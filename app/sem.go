package app

import (
	"fmt"
	//"time"
)

var sem = make(chan int, 2)   //max outgoing is 2


func process(r string){
	//time.Sleep(100 * time.Millisecond)
	fmt.Println(r)
}

func Serve(queue chan string) {
	for req :=range queue {
		req := req // !!!
		sem <- 1
		go func() {
			process (req)
			<-sem
		}()
	}
}
