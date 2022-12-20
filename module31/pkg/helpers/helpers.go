package helpers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func RemoveElementFromArray(friends []string, friendName string) []string {
	needIndex := 0
	for index, friend := range friends {
		if friend == friendName {
			needIndex = index
			break
		}
	}
	newFriends := append(friends[:needIndex], friends[needIndex+1:]...)
	return newFriends
}

func SaveFile(data, filename string) error {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("Не удалось открыть файл: %s", err.Error())
	}
	defer f.Close()
	var buffer bytes.Buffer
	buffer.WriteString(data)
	_, err = f.Write(buffer.Bytes())
	if err != nil {
		panic(err)
	}
	return nil
}

func ReadFile(filename string) ([]byte, error) {
	fmt.Println(filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Не удалось открыть файл: %s", err.Error())
	}

	return data, nil
}
