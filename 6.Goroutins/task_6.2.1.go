package main

import (
	"fmt"
	"sync"
)

func factorial(a int, ch chan int) {
    result := 1
    for i := 1; i <= a; i++ {
        result = result * i
    }
    ch <- result
}
func main(){
		var a, b int
		fmt.Scan(&a, &b)
    wg := new(sync.WaitGroup)
    ch1 := make(chan int)
    ch2 := make(chan int)
    
    wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
            factorial(a, ch1)
            factorial(b, ch2)
		}(wg)
    // wg.Wait()
    
    res1 := <-ch1
    res2 := <-ch2
    fmt.Println(res1 + res2)
}