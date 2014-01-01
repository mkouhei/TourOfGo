//http://go-tour-jp.appspot.com/#11
package main

import (
	"fmt"
)

var x, y, z int
var c, python, java bool

func main() {
	fmt.Println(x, y, z, c, python, java)

	x, y, z, c, python, java := "hoge", "moge", true, 1, "foo", 0.0
	fmt.Println(x, y, z, c, python, java)
}
