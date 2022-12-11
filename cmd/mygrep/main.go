package main

import (
	// Uncomment this to pass the first stage
	"bytes"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// Usage: echo <input_text> | your_grep.sh -E <pattern>
func main() {
	//fmt.Println("My arguements are ", os.Args)
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	// myString := string(line)
	// fmt.Println("This is my line ", myString)
	// fmt.Println("This is my pattern ", pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		//fmt.Println("Pattern didn't match")
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern string) (bool, error) {
	//fmt.Println("This is my pattern ", pattern)
	// To see if pattern size is only equal to 1

	if utf8.RuneCountInString(pattern) != 1 {
		if pattern == "\\d" {
			//var isdigit = false
			//fmt.Println("This is my pattern /d")
			for i := 0; i < len(line); i++ {
				if unicode.IsNumber(rune(line[i])) {
					//fmt.Println("number hai toh chalega bhidu")
					return true, nil

				}
			}
		}
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	var ok bool

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	//fmt.Println("Logs from your program will appear here!")

	// Uncomment this to pass the first stage
	ok = bytes.ContainsAny(line, pattern)

	return ok, nil
}
