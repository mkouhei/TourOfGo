//http://go-tour-jp.appspot.com/#5
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.",
		math.Nextafter(2, 3))
}
