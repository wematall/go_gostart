package main

// ФАКТОРИАЛ ЧИСЛА

// Напишите программу на Go,
// которая запрашивает у пользователя положительное целое число n
// и вычисляет его факториал.
// Факториал числа n обозначается как n!
// и равен произведению всех целых чисел от 1 до n.

import "fmt"

func factorial(a int) int{
	res := 1
	for i := 1; i <=a; i++ {
		res *= i
	}
	return res
}

func main() {
	var a int

	fmt.Println("Вычислим факториал")
	fmt.Print("Введите число: ")
	
	fmt.Scan(&a)

	res := factorial(a)
	fmt.Println(res)
}