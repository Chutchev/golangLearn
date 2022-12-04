package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	// var fileName1, fileName2, fileName3 string
	switch len(os.Args) - 1 {
	case 0:
		log.Fatal("Не было передано ни одно аргумента")
	case 1:
		{
			result, err := oneFileAction(os.Args[1])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(result)
		}
	case 2:
		{
			result, err := twoFilesAction(os.Args[1], os.Args[2])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(result))
		}
	case 3:
		{
			result, err := threeFilesAction(os.Args[1], os.Args[2], os.Args[3])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(result))
		}
	}

}

func oneFileAction(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("Файл %v не найден!", filename)
	}
	defer f.Close()
	resultBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("Не удалось прочитать файл!")
	}
	return string(resultBytes), nil

}

func twoFilesAction(filename1, filename2 string) (string, error) {
	resultArray := make([]string, 0)
	result1, err := oneFileAction(filename1)
	if err != nil {
		return "", err
	}
	resultArray = append(resultArray, result1)
	result2, err := oneFileAction(filename2)
	if err != nil {
		return "", err
	}
	resultArray = append(resultArray, result2)
	return strings.Join(resultArray, "\n"), nil
}

func threeFilesAction(filename1, filename2, fileName3 string) (string, error) {
	resultArray, err := twoFilesAction(filename1, filename2)
	if err != nil {
		return "", err
	}
	f, err := os.Create(fileName3)
	if err != nil {
		return "", fmt.Errorf("Не удалось создать файл %v", fileName3)
	}
	defer f.Close()
	resultArrayBytes := []byte(resultArray)
	ioutil.WriteFile(fileName3, resultArrayBytes, 0666)
	return fmt.Sprintf("Объединение прошло успешно"), nil
}
