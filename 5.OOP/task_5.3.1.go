package main

// Напишите функцию Swap,
// которая принимает два указателя на целочисленные переменные
// и меняет их значения местами.

import "fmt"

func Swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 10
	b := 20

	fmt.Println("До Swap: a =", a, "b =", b)

	Swap(&a, &b)

	fmt.Println("После Swap: a =", a, "b =", b)
}