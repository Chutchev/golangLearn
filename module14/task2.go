package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generate2dPoint(a int) (int, int) {
	return rand.Intn(a), rand.Intn(a)
}

func newPoint(x, y int) (newX, newY int) {
	newX = 2*x + 10
	newY = -3*y - 5
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var max int
	fmt.Println("Введите максимальную цифру для генерации числа: ")
	fmt.Scan(&max)
	x1, y1 := generate2dPoint(max)
	x2, y2 := generate2dPoint(max)
	x3, y3 := generate2dPoint(max)

	fmt.Printf("Изначальные точки:\n\t1.(%d, %d)\n\t2.(%d, %d)\n\t3.(%d, %d)\n", x1, y1, x2, y2, x3, y3)

	newX1, newY1 := newPoint(x1, y1)
	newX2, newY2 := newPoint(x2, y2)
	newX3, newY3 := newPoint(x3, y3)

	fmt.Printf("Новые точки точки:\n\t1.(%d, %d)\n\t2.(%d, %d)\n\t3.(%d, %d)\n", newX1, newY1, newX2, newY2, newX3, newY3)
}
