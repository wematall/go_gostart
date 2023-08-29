package main

// Напишите функцию "averageValue",
// которая принимает три целых числа
// в качестве входных данных
// и возвращает их среднее значение, округленное до целого числа,
// при условии, что округление происходит в меньшую сторону.

import "fmt"

func averageValue(a, b, c int) int {
	return (a+b+c)/3
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	result := averageValue(a, b, c)
	fmt.Println(result)
}