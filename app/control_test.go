package app

import (
	"testing"
	"fmt"
)

func TestSwitch(t *testing.T){
	s:="5G-AKA2"
	switch s {
	default:
		fmt.Println("default")
	case "5G-AKA":
		fmt.Println("birdy")
	}
}