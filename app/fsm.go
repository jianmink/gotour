package app

import (
	"github.com/looplab/fsm"
	"fmt"
)

func beforeEventAny(event *fsm.Event){
	fmt.Printf("callback invokded before Event Any")
}

func beforeEventOpen(event *fsm.Event){
	fmt.Printf("callback invokded before Event (%v) \n",event.Event)
}



func NewMyFSM() *fsm.FSM {
	myfsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"before_event": beforeEventAny,
			"before_open": beforeEventOpen,
		},
	)

	return myfsm
}