package main

// Напишите код с помощью структуры,
// чтобы был вывод как в тестовых данных.
// Martin 35 male

import "fmt"

type my_struct struct {
	name string
	age int
	gender string
}

func main() {
	st := my_struct{"Martin", 35, "male"}
	fmt.Println(st.name, st.age, st.gender)
}