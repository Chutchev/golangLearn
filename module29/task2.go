package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	c := make(chan os.Signal, 1)
	numberChan := make(chan string)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	for {
		go square(numberChan)
		select {
		case b := <-numberChan:
			fmt.Println(b)
		case <-c:
			closeAll()
			os.Exit(1)
		}
	}
}

func square(numberChan chan string) {
	var number int
	fmt.Println("Введите число: ")
	fmt.Scan(&number)
	numberChan <- fmt.Sprintf("%d", number*number)
}

func closeAll() {
	fmt.Println("Close all!")
}
