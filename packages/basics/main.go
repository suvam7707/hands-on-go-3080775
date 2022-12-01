// packages/basics/main.go
package main

import (
	"fmt"
)

var val = "global"

func getGlobal() {
	fmt.Println("global val =", val)
	val = "updated global"
	fmt.Println("global val =", val)
}

func main() {
	val := 42
	fmt.Printf("%T, local val = %v\n", val, val)
	getGlobal()
	loc := &val
	fmt.Printf("%T, local &val = %v\n", loc, loc)
	*loc = 100
	fmt.Println("local val =", val)
}
