package main

import (
	"flag"
	"fmt"
)

func main() {
	var str, substr string
	flag.StringVar(&str, "str", "Строка в которой необходимо проверить", "Строка в которой необходимо проверить")
	flag.StringVar(&substr, "substr", "строка для поиска", "строка для поиска")
	flag.Parse()

	fmt.Println(str, substr)
	fmt.Println(findInStroke(str, substr))
}

func findInStroke(str, substr string) bool {
	runesStr := []rune(str)
	runesSubStr := []rune(substr)
	isReturn := true
	for r := 0; r < len(runesStr); r++ {
		isBreak := false
		for rs := 0; rs < len(runesSubStr); rs++ {
			isReturn = true
			newIndex := r + rs
			if newIndex >= len(runesStr) {
				isBreak = true
				break
			}
			if runesStr[r+rs] != runesSubStr[rs] {
				isReturn = false
			}
		}
		if isBreak {
			break
		}
		if isReturn {
			return true
		}
	}
	return false
}
