//http://go-tour-jp.appspot.com/#56

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
	//return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}

func Sqrt(x float64) (float64, error) {

	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}

	x64 := float64(x)
	xn := float64(0)
	n := 10
        
	for ; xn * xn < 64; xn++ {
	}

	for i := 1; i < n; i++ {
		xn = (xn + x64 / xn) / 2
	}
	return xn, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
