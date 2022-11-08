package main

import (
	"fmt"
	"strings"
)

func main() {
	stroke := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software."
	count := 0
	for len(stroke) > 0 {
		spaceIndex := strings.Index(stroke, " ")
		if spaceIndex > 0 {
			word := stroke[:spaceIndex]
			// fmt.Println(word)
			isEquals := strings.Compare(word, strings.Title(word))
			if isEquals == 0 {
				count++
			}
			stroke = stroke[spaceIndex+1:]
			// fmt.Println(stroke)
		} else {
			isEquals := strings.Compare(stroke, strings.Title(stroke))
			if isEquals == 0 {
				count++
			}
			break
		}
	}
	fmt.Println(count)
}
