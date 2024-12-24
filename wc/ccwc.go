package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func c() int {
	fileName := flag.String("c", "test.txt", "File name parameter")
	length := 0
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	// Create a reader
	reader := bufio.NewReader(file)

	// Count the number of bytes
	for {
		_, err := reader.ReadByte()
		if err != nil {
			break // End of file
		}
		length += 1
	}
	return length
}

func l() int {
	fileName := flag.String("l", "test.txt", "File name parameter")
	length := 0
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		length += 1
	}
	return length
}

func w() int {
	fileName := flag.String("w", "test.txt", "File name parameter")
	length := 0
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		length += len(strings.Fields(reader.Text()))
	}
	return length
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Parameters are needed")
	}
	fileName := new(string)
	*fileName = "test.txt"
	if os.Args[1] == "-c" {
		fmt.Printf("%d %s\n", c(), *fileName)
		
	} else if os.Args[1] == "-l" {
		fmt.Printf("%d %s\n", l(), *fileName)
		
	} else if os.Args[1] == "-w" {
		fmt.Printf("%d %s\n", w(), *fileName)
		
	} else {
		fmt.Printf("%d %d %d %s\n", l(), w(), c(), *fileName)
	}
}