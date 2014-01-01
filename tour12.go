//http://go-tour-jp.appspot.com/#12
package main

import (
	"fmt"
)

var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

func main() {
	fmt.Println(x, y, z, c, python, java)

	x, y, z, c, python, java := "hoge", "moge", true, 1, "foo", 0.0
	//x, y, z, c, python, java = "hoge", "moge", true, 1, "foo", 0.0
	fmt.Println(x, y, z, c, python, java)
}
