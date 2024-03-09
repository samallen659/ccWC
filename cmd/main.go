package main

import (
	"flag"
	"fmt"
)

func main() {
	var byteFlag bool
	flag.BoolVar(&byteFlag, "c", false, "Show byte count of input")
	flag.Parse()

	if byteFlag {
		fmt.Println("true")
	}
}
