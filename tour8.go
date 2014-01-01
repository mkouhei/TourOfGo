//http://go-tour-jp.appspot.com/#8
package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Happy new year", add(2013, 1))
}
