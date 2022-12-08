package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	var number int
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	channel := make(chan int)
	for {
		fmt.Println("Введите число!")
		fmt.Scan(&number)
		select {
		default:
			fmt.Println(<-square(number, channel))
		case <-c:
			os.Exit(1)
		}
	}

}

func square(number int, channel chan int) chan int {
	result := number * number
	channel <- result
	return channel
}
