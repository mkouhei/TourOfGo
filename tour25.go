//http://go-tour-jp.appspot.com/#25

package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
	fmt.Println(Vertex{1.0, 2.0})
}
