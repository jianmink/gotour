package main

import (
	"fmt"
	"github.com/jianmink/gotour/app"
	"log"
	"os"
	"time"
)

func Sum(a, b int) (int) {
	return a + b
}


func embed() (int, error){

	type Job struct {
		Command string
		*log.Logger
	}

	job := &Job{"hello", log.New(os.Stderr, "Job: ", log.Ldate)}
	job.Println("doom")

	return 0, nil
}

func main() {
	fmt.Println("hello world")

	embed()
	app.WaitResponsePlus2()
	//app.WaitResponse2()

	q := make(chan string, 3) // 3 requests
	q <- "three"
	q <- "two"
	q <- "one"

	go app.Serve(q)

	time.Sleep(1 * time.Second)
	close(q)
}
