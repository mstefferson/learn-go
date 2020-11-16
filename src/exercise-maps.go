package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	wc := make(map[string]int)
	var v int
	var ok bool
	for index, word := range strings.Fields(s) {
		fmt.Printf("%v %v", index, word)
		v, ok = wc[word]
		if ok{
			wc[word] = v + 1
		}else{
			wc[word] = 1
		}
	}
	return wc
}

func main() {
	wc.Test(WordCount)
}
