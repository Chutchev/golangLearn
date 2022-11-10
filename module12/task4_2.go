package main

import (
	"fmt"
	"io/ioutil"
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
	resultBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("не удалось прочитать файл")
	}
	fmt.Println(string(resultBytes))
}
