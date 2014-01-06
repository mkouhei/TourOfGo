//http://go-tour-jp.appspot.com/#44

package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.


func fibonacci() func() int {
	fn, fn1 := -1, 0
	return func() int {
		if fn < 1 {
			fn++
			return fn
		}
		fn1, fn = fn, fn1 + fn
		return fn
	}
}

func main() {
	f := fibonacci()
	for i:= 0; i < 10; i++ {
		fmt.Println(f())
	}
}

