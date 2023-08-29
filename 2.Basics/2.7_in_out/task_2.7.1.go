package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Напишите программу на Go,
// которая запрашивает у пользователя его имя и возраст,
// а затем выводит приветственное
// сообщение в следующем формате: "Привет, [имя]! Тебе [возраст] лет."

func main() {
    reader := bufio.NewReader(os.Stdin)
    var name string
    var age string
	fmt.Print("Введите ваше имя: ")
    name, _ = reader.ReadString('\n')
    name = strings.Trim(name, "\n")
	fmt.Print("Введите ваш возраст: ")
    age, _ = reader.ReadString('\n')
    age = strings.Trim(age, "\n")
	fmt.Printf("Привет, %s! Тебе %s лет.\n", name, age)
}