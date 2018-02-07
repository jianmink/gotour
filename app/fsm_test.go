package app

import (
	"testing"
	"fmt"
)

func TestNewMyFSM(t *testing.T) {
	fsm := NewMyFSM()
	fmt.Println(fsm.Current())
	fsm.Event("open")
}
