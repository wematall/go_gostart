package main

import "fmt"

// Зачетная книжка
// Напишите функцию, которая находит среднее значение из трех переданных ей баллов
// и возвращает буквенное значение,
// связанное с этой оценкой.(используйте switch-case)

// Числовой счет	Оценка по буквам
// 90 <= оценка <= 100	'A'
// 80 <= оценка < 90	'B'
// 70 <= оценка < 80	'C'
// 60 <= оценка < 70	'D'
// 0 <= оценка < 60	'F'

func GetGrade(a, b, c int) string {
	sum := (a + b + c)/3

	switch {
	case sum >= 90 && sum <= 100:
		return "A"
	case sum >= 80 && sum < 90:
		return "B"
	case sum >= 70 && sum < 80:
		return "C"
	case sum >= 60 && sum < 70:
		return "D"
	case sum >= 0 && sum < 60:
		return "F"
	default:
		return "undefined mark"
	}
}

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	fmt.Println(GetGrade(a, b, c))
}