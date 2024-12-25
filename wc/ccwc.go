package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func c(fileName string) int {
	length := 0
	file, err := os.Open(fileName)
	if err != nil {log.Fatal(err)}

	reader := bufio.NewReader(file)

	for {
		_, err := reader.ReadByte()
		if err != nil {
			break // End of file
		}
		length += 1
	}
	return length
}

func l(fileName string) int {
	length := 0
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		length += 1
	}
	return length
}

func w(fileName string) int {
	length := 0
	file, err := os.Open(fileName)
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
	if os.Args[1] == "-c" {
		fmt.Printf("%d %s\n", c(os.Args[2]), os.Args[2])
		
	} else if os.Args[1] == "-l" {
		fmt.Printf("%d %s\n", l(os.Args[2]), os.Args[2])
		
	} else if os.Args[1] == "-w" {
		fmt.Printf("%d %s\n", w(os.Args[2]), os.Args[2])
		
	} else {
		fmt.Printf("%d %d %d %s\n", l(os.Args[1]), w(os.Args[1]), c(os.Args[1]), os.Args[1])
	}
}