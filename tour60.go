//http://go-tour-jp.appspot.com/#60
//https://gist.github.com/zyxar/2317744#file-exercise-tour-go-L154
//cd $(git rev-parse --show-toplevel)
//export GOPATH=$(pwd)
//go build tour60.go
//base64 -d <(./tour60 | sed 's/^IMAGE://') > tour60.png

package main

import (
	"go-tour-72382f964b32/pic"
	"image"
	"image/color"
)

type Image struct{
	Width, Height int
	colour uint8
}

func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}

func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colour+uint8(x), r.colour+uint8(y), 255, 255}
}

func main() {
	m := Image{100, 100, 128}
	pic.ShowImage(&m)
}
