package app

import (
	"regexp"
	"testing"
	"fmt"
)

func TestRegMatch(t *testing.T){
	s := "\"LinkToConfirm\":\"20180306161415\""
	p := `\"LinkToConfirm\":\"[0-9]+\"`

	r, err := regexp.Match(p,[]byte(s))
	fmt.Printf("r: %v, err: %v", r, err)

}