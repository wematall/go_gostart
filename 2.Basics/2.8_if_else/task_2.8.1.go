package main

import "fmt"

func maxThree(a, b, c int) int{
	if a > b && a > c {
		return a
	} else if b > a && b > c {
		return b
	} else {
		return c
	}

}

func main(){
    var (
		a int
		b int
		c int
	)
	fmt.Scan(&a, &b, &c)
    average:=maxThree(a,b,c)
    fmt.Println(average)
}