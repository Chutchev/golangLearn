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
	reverseSorted := func(array []int) ([]int, error) {
		if len(array) == 0 {
			return make([]int, 0), fmt.Errorf("Передано 0 элементов массива")
		}
		for i := 0; i < len(array)-1; i++ {
			for j := i; j < len(array); j++ {
				if array[j] > array[i] {
					array[i], array[j] = array[j], array[i]
				}
			}
		}
		return array, nil
	}
	newArray, err := reverseSorted(array)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(newArray)
}

func fillArray(array []int) []int {
	for i := 0; i < SIZE; i++ {
		array = append(array, rand.Intn(SIZE*SIZE))
	}
	return array
}
