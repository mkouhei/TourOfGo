//http://go-tour-jp.appspot.com/#24
//copy from https://gist.github.com/yuya-takeyama/6662167

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	x64 := float64(x)
	xn := float64(0)
	n := 10
	
	for ; xn * xn < 64; xn++ {
	}

	for i := 1; i < n; i++ {
		xn = (xn + x64 / xn) / 2
	}
	return xn
}

func main() {
	fmt.Println(Sqrt(1), Sqrt(2), Sqrt(3), Sqrt(1001))
	fmt.Println(math.Sqrt(1), math.Sqrt(2), math.Sqrt(3), math.Sqrt(1001))
}
