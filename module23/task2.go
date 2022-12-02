package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	sentences := [4]string{"Hello world", "Hello skillSbox", "Привет Мир", "Привет skillSbox"}
	chars := [5]rune{'H', 'E', 'L', 'S', 'W'}

	for i := range sentences {
		sentences[i] = strings.ToUpper(sentences[i])
	}

	symbolsInStrokes, err := findSymbolsInStrokes(sentences[:], chars[:])
	if err != nil {
		log.Fatalf("ERROR YOPTA: %v", err)
	}

	for i := range symbolsInStrokes {
		fmt.Println(symbolsInStrokes[i])
	}
}

func findSymbolsInStrokes(sentences []string, symbols []rune) ([4][5]string, error) {

	var result [4][5]string

	if len(sentences) == 0 || len(symbols) == 0 {
		return result, fmt.Errorf("Are you AHUELI tam? Длина одного или сразу двух массивов равны 0!!!")
	}

	for j := range sentences {
		for i := range symbols {
			for k := len(sentences[j]) - 1; k >= 0; k-- {
				if symbols[i] == rune(sentences[j][k]) {
					fmt.Printf("%v, postion %v\n", string(symbols[i]), k)
					result[j][i] = fmt.Sprintf("%v, postion %v\n", string(symbols[i]), k)
					break
				}
			}
		}
	}
	return result, nil
}
