//http://go-tour-jp.appspot.com/#48
// copy from https://github.com/fgrehm/go-tour/blob/master/48-complex-cube-roots.go

package main

import (
	"fmt"
	"math/cmplx"
)

const delta = 1e-15

func Cbrt(x complex128) complex128 {
	oldz, z := complex128(0.0), complex128(1.0)
	for {
		z = z - (z * z * z - x) / (3 * z * z)
		if cmplx.Abs(z - oldz) <= delta {
			break
		}
		oldz = z
	}
	return z
	
}

func main() {
	fmt.Println(Cbrt(2))
	fmt.Println(cmplx.Pow(2, 1.0/3))
}
