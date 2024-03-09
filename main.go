package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type FileDetails struct {
	byteCount int
	lineCount int
	wordCount int
	charCount int
}

func main() {
	var byteFlag bool
	var lineFlag bool
	var wordFlag bool
	var charFlag bool
	flag.BoolVar(&lineFlag, "l", false, "Show line count of input")
	flag.BoolVar(&byteFlag, "c", false, "Show byte count of input")
	flag.BoolVar(&wordFlag, "w", false, "Show word count of input")
	flag.BoolVar(&charFlag, "m", false, "Show character count of input")
	flag.Parse()

	//flag.Args() will be greater than one if the flags are specified after the file name
	//this will cause the flag to have not been parsed correctly
	if len(flag.Args()) > 1 {
		log.Fatal("Invalid input")
	}

	fileName := flag.Args()[0]

	var fileDetails *FileDetails
	var err error
	// if no flags set, true for all
	if !byteFlag && !lineFlag && !wordFlag && !charFlag {
		fileDetails, err = CalculateFileDetails(fileName, true, true, true, true)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(fileDetails)
}

func CalculateFileDetails(fileName string, c bool, l bool, w bool, m bool) (*FileDetails, error) {
	fileDetails := &FileDetails{}

	fileBytes, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	if c {
		fileDetails.byteCount = len(fileBytes)
	}

	return fileDetails, nil
}
