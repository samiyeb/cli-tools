package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	
	if len(os.Args) == 1 {
		fmt.Println("Invalid input")
		return
	}

	var file *os.File
	delim := `\t`

	for i := 1; i < len(os.Args); i++ {
		if tempFile, err := os.Open(os.Args[i]); err == nil {
			file = tempFile
		} else if len(os.Args[i]) >= 2 && os.Args[i][:2] == `-d`{
			delim = string(os.Args[i][2])
		}
	}

	var scanner *bufio.Scanner

	for i := 0; i < len(os.Args); i++ {
		scanner = bufio.NewScanner(file)
		if len(os.Args[i]) >= 2 && os.Args[i][:2] == `-f` {
			if _, err := strconv.Atoi(string(os.Args[i][2])); err == nil{
				for scanner.Scan() {
					var wordsList []string
					if delim == `\t` {
						wordsList = strings.Fields(scanner.Text())
					} else {
						wordsList = strings.Split(scanner.Text(), delim)
					}
					
					colListSpace := strings.Fields(string(os.Args[i][2:]))
					colListComma := strings.Split(string(os.Args[i][2:]), ",")

					var colList []string 

					if len(colListSpace) > len(colListComma) {
						colList = colListSpace
					} else {
						colList = colListComma
					}

					for i := 0; i < len(colList); i++ {
						col, err := strconv.Atoi(string(colList[i]))
						if err != nil {
							fmt.Println("Error in column list for -f")
						}
						fmt.Printf("%s\t", wordsList[col - 1])
					}
					fmt.Printf("\n")
				}
			}
		}
	}


}