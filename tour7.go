//http://go-tour-jp.appspot.com/#7
package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println("Happy new year", add(2013, 1))
}
