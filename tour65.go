//http://go-tour-jp.appspot.com/#65

package main

import (
	"fmt"
)


func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	//c <- 3 // dead blocks +0x9a
	fmt.Println(<-c)
	fmt.Println(<-c)
	//fmt.Println(<-c) // dead block +0x195
}
