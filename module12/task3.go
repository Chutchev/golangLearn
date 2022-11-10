package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "task3.txt"
	fmt.Printf("Создаю файл %s\n", fileName)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Не удалось создать файл")
		return
	}
	defer file.Close()
	fmt.Println("Меняю доступы к файлу!")
	file.Chmod(0222)
	fmt.Println("Закрываю файл!")
	file.Close()
	fmt.Printf("Открываю файл: %s снова\n", fileName)
	newFile, err := os.Open(fileName)
	if err == nil {
		fmt.Println("Файл почему то открылся!")
		return
	}
	fmt.Println("Не удалось открыть файл")
	defer newFile.Close()

}
