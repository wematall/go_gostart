package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// найти самое короткое слово в строке и
// вывести длину короткого слова

func FindShort(s string) int {
	my_slice := strings.Split(s, " ")
	result := len(my_slice[0])

	for _, v := range my_slice {
		if result > len(v) {
			result = len(v)
		}
	}

	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	fmt.Println(FindShort(str))
}