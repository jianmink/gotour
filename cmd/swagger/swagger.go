package main

import (
	"os"
	"fmt"
	"path/filepath"
	"github.com/jianmink/gotour/openapi"
)

func main() {
	if len(os.Args) < 1 {
		fmt.Println("No spec file (json/yaml)")
	} else {
		f := os.Args[1]
		ext := filepath.Ext(f)
		f1 := filepath.Base(f)
		d := f1[0:len(f1)-len(ext)] + ".go"
		fmt.Printf("input: %v, output: %v\n", f1, d)
		openapi.DecodeSpecFile(f,d)
	}
}

