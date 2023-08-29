package main

// В этой задаче вы должны будете написать пару простых функций:

// Функция SumFloat - функция которая принимает на себя два числа типа float32 и возвращает их сумму.
// Функция MultiplyFloat -возвращает их произведение.
// Функция PowFloat - функция которая принимает на себя два числа a и b (float64)
// возвращает число a в степени b(a^b).
// В данной функции можете воспользоваться функцией из пакета Math

import (
	"fmt"
	"math"
)

func SumFloat(a, b float32) float32{
	return a+b
}

func MultiplyFloat(a, b float32) float32{
	return a*b
}

func PowFloat(a, b float64) float64{
	return math.Pow(a, b)
}

func main(){
var (
	a float32
	b float32
)
fmt.Scan(&a, &b)
	sumFloat:=SumFloat(a,b)
	multiplyFloat:=MultiplyFloat(a,b)
	powFloat:=PowFloat(float64(a),float64(b))
	output := fmt.Sprintf("%.1f %.1f %.3f", sumFloat, multiplyFloat, powFloat)
	fmt.Println(output)
}