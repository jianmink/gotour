package app

import (
	"github.com/looplab/fsm"
	"fmt"
)

func FsmCan() {
	myfsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(myfsm.Can("open"))
	fmt.Println(myfsm.Can("close"))
}