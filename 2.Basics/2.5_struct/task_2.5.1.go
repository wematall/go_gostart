package main

import "fmt"

// написать структуру

type Person struct {
    id int
    Name string
    Surname string
}

func main(){
    var (
		id int
		name string
		sname string
	)
	fmt.Scan(&id, &name, &sname)
	prs := Person{id, name, sname}
	fmt.Println(prs)
}