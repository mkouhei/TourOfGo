//http://go-tour-jp.appspot.com/#36
//http://code.google.com/p/go-tour/source/browse/pic/pic.go
//wget http://go-tour.googlecode.com/archive/72382f964b32d8c3863aa4e4a5852b58a771e22e.zip
//unzip 72382f964b32d8c3863aa4e4a5852b58a771e22e.zip
//export GOPATH=$(pwd)
//go build tour36.go
//base64 -d <(./tour36 | sed 's/^IMAGE://') > tour64.png

package main

//import "code.google.com/p/go-tour/pic"
import "go-tour-72382f964b32/pic"

func Pic(dx, dy int) [][]uint8 {
	var p = make([]([]uint8), dy)
	for i := 0; i < len(p); i++ {
		p[i] = make([]uint8, dx)
		for j := 0; j < len(p[i]); j++ {
			//p[i][j] = uint8((i + j) / 2)
			//p[i][j] = uint8(i * j)
			p[i][j] = uint8(i ^ j)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}

