package main

import (
	"fmt"
	"operation"
)

func main() {
	var fPath string
	fmt.Scan(&fPath)
	average, _ := operation.Average(fPath)
	fmt.Println(average)
	sum, _ := operation.Sum(fPath)
	fmt.Println(sum)
}
