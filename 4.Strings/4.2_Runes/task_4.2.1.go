package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Напишите функцию accum, которая принимает строку
// в качестве аргумента и возвращает новую строку,
// в которой каждый символ повторяется
// в зависимости от его позиции в строке.
// При этом каждый повтор символа должен быть написан с учетом регистра:
// первая буква - в верхнем регистре, остальные - в нижнем регистре.
// Символы в новой строке должны быть разделены дефисом (-).

// accum("abcd") -> "A-Bb-Ccc-Dddd"
// accum("RqaEzty") -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
// accum("cwAt") -> "C-Ww-Aaa-Tttt"

func accum(s string) string {
	var result string
	my_slice := []rune(s)
	for i, v := range my_slice {
		result = result + strings.ToUpper(string(v))
		for k := 1; k <= i; k++ {
			result = result + string(v)
		}

		if i == len(my_slice) -1 {
			break
		}
		result = result + "-"
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.Trim(str, "\n")
	str = strings.ToLower(str)

	fmt.Println(accum(str))
}