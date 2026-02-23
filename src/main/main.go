package main

import (
	"fmt"
	"operation"
	"readfile"
)

func main() {
	data, err := readfile.ReadFile("text.txt")
	if err != nil {
		fmt.Println("Ошибка!", err)
	}
	res, err := operation.Sum(data)
	if err != nil {
		fmt.Println("Ошибка!", err)
	}
	fmt.Println("Результат расчета:", res)
}
