//http://go-tour-jp.appspot.com/#35

package main

import (
	"fmt"
)

func main() {
	pow := make([]int, 10)
	for i:= range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
