//http://go-tour-jp.appspot.com/#41
//see build procedure comment of tour36.go

package main

//import "code.google.com/p/go-tour/wc"
import (
	"strings"
	"go-tour-72382f964b32/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	wordCount := make(map[string]int)

	var ok bool
	for i := 0; i < len(words); i++ {
		_, ok = wordCount[words[i]]
		if ok == true {
			wordCount[words[i]] += 1
		} else {
			wordCount[words[i]] = 1
		}
	}
	//return map[string]int{"x": 1}
	return wordCount
}

func main() {
	wc.Test(WordCount)
}

