package main

import (
	"flag"
	"fmt"
)

func main() {
	var byteFlag bool
	var lineFlag bool
	var fileName string
	flag.BoolVar(&lineFlag, "l", false, "Show line count of input")
	flag.BoolVar(&byteFlag, "c", false, "Show byte count of input")
	flag.Parse()
	fmt.Println(fileName)
	fmt.Println(flag.Args())

	if byteFlag {
		fmt.Println("true")
	}
}
