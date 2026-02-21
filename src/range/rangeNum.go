package rangeNum

import "fmt"

// Функция RangeNum позволяет получить промежуток с файла
func RangeNum() (fNum, sNum int) {
	// Первое число из массива
	fmt.Println("Введите первое значение:")
	fmt.Scan(&fNum)
	fmt.Println("Введите второе значение:")
	fmt.Scan(&sNum)
}
