//http://go-tour-jp.appspot.com/#33

package main

import (
	"fmt"
)

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z))

	if z == nil {
		fmt.Println("nil")
	}
}
