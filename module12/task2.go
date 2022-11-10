package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "posts.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Файла не существует!", err)
		return
	}
	defer file.Close()
	fInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}
	size := fInfo.Size()
	if size == 0 {
		fmt.Println("Файл пуст")
		return
	}
	buf := make([]byte, size)
	if _, err := io.ReadFull(file, buf); err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
}
