// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	//
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from panic:", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	//
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("failed to read file: %s", err))
	}

	// convert the bytes to a string
	//
	str := string(data)

	// initialize a map to store the counts
	//
	m := map[rune]int{}

	// use the standard library utility package that can help us split the string into words
	//
	strs := strings.Split(str, " ")

	// capture the length of the words slice
	//
	l := len(strs)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	//
	for i := 0; i < l; i++ {
		for _, ch := range strs[i] {
			_, exists := m[ch]
			if !exists {
				m[ch] = 1
			} else {
				m[ch]++
			}
		}
	}

	// dump the map to the console using the spew package
	//
	spew.Dump(m)
}
