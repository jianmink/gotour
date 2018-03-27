package main

import (
	"os"
	"fmt"
	_ "path/filepath"
	"github.com/jianmink/gotour/openapi"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("swagger <spec file> <dest dir> <service> <version>")
		fmt.Println("example: swagger ausf.json ./output ausf v1")
	} else {
		//f := os.Args[1]
		//ext := filepath.Ext(f)
		//f1 := filepath.Base(f)
		//d := f1[0:len(f1)-len(ext)] + ".go"

		spec := os.Args[1]
		dst := os.Args[2]
		service := os.Args[3]
		version := os.Args[4]
		fmt.Printf("input: %v, output: %v\n", spec, dst+"/"+service+version+"/struct.go")

		openapi.DecodeSpecFile(spec, dst, service, version)
	}
}
