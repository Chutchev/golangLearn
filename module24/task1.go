package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const SIZE = 10

func main() {
	rand.Seed(time.Now().Unix())
	array := make([]int, 0, SIZE)
	array = fillArray(array)
	fmt.Println(array)
	array, err := insertedSort(array)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(array)
}

func fillArray(array []int) []int {
	for i := 0; i < SIZE; i++ {
		array = append(array, rand.Intn(SIZE*SIZE))
	}
	return array
}

func insertedSort(array []int) ([]int, error) {
	if len(array) == 0 {
		return make([]int, 0), fmt.Errorf("Передано 0 элементов массива")
	}
	for i := 1; i < len(array); i++ {
		value := array[i]
		needIndex := i
		for ; needIndex >= 1 && array[needIndex-1] > value; needIndex-- {
			array[needIndex] = array[needIndex-1]
		}
		array[needIndex] = value
	}
	return array, nil
}
