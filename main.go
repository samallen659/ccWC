package main

import (
	"bytes"
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
	// if no flags set, true for all but charCount
	if !byteFlag && !lineFlag && !wordFlag && !charFlag {
		fileDetails, err = CalculateFileDetails(fileName, true, true, true, false)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("    %d  %d  %d %s\n", fileDetails.lineCount, fileDetails.wordCount, fileDetails.byteCount, fileName)
		return
	}

	fileDetails, err = CalculateFileDetails(fileName, byteFlag, lineFlag, wordFlag, charFlag)
	if err != nil {
		log.Fatal(err)
	}

	output := "   "
	if lineFlag {
		output = fmt.Sprintf("%s  %d", output, fileDetails.lineCount)
	}
	if wordFlag {
		output = fmt.Sprintf("%s  %d", output, fileDetails.wordCount)
	}
	if byteFlag {
		output = fmt.Sprintf("%s  %d", output, fileDetails.byteCount)
	}
	if charFlag {
		output = fmt.Sprintf("%s  %d", output, fileDetails.charCount)
	}

	output = fmt.Sprintf("%s %s", output, fileName)

	fmt.Println(output)
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

	if l {
		fileDetails.lineCount = bytes.Count(fileBytes, []byte{10})
	}

	if w {
		count := 0
		lines := bytes.Split(fileBytes, []byte{10})
		//exclude last entry in lines due to splitting on newline creating empty slice at the end
		for _, line := range lines[:len(lines)-1] {
			line = bytes.Trim(line, " ")
			line = bytes.Trim(line, "\t")
			if len(line) == 0 {
				continue
			}

			spaceSplit := bytes.Split(line, []byte{32})
			for _, words := range spaceSplit {
				tabSplit := bytes.Split(words, []byte{9})
				count += len(tabSplit)
			}
		}

		fileDetails.wordCount = count
	}

	if m {
		fileDetails.charCount = len([]rune(string(fileBytes)))
	}

	return fileDetails, nil
}
