package app


type Request struct {
	args        []int
	f           func([]int) int
	resultChan  chan int
}


func sum (a []int) (s int){
	for _, v := range a {
		s += v
	}
	return
}

func handle(queue chan *Request) {
	for req := range queue {
		req.resultChan <- req.f(req.args)
	}
}


