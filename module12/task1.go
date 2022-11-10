package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fileName := "posts.txt"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var buffer bytes.Buffer
	num := 0
	for {
		num++
		fmt.Println("Введите текст")
		scanner.Scan()

		text := scanner.Text()
		if strings.ToLower(text) == "exit" {
			break
		}
		current_datetime := time.Now()
		format_current_datetime := current_datetime.Format("2006-01-02 15:04:05")
		buffer.WriteString(fmt.Sprintf("%d. %s %s\n", num, format_current_datetime, text))
	}
	_, err = file.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	fmt.Println("Вышли из программы")
}
