package main

import "fmt"

const var1 = 150
const var2 = 22
const var3 = 7

func addVar1(n int) int {
	return n + var1
}

func addVar2(n int) int {
	return n + var2
}

func addVar3(n int) int {
	return n + var3
}

func main() {
	var a int
	fmt.Println("Введите число")
	fmt.Scan(&a)
	d := addVar3(addVar2(addVar1(a)))
	fmt.Println(d)
}
