package operation

import (
	"readfile"
)

// Average - вычисляет среднее значение из тектового файла
func Average(fPath string) (average float64, err error) {
	// Считывает значение из текстового файла
	array, err := readfile.ReadFile(fPath)
	// Проверка на ошибки
	if err != nil {
		return average, err
	}
	// Суммирует полученные значения из файла функцией Sum
	sum, err := Sum(fPath)
	if err != nil {
		return average, err
	}
	// Считывает длину массива
	arrLen := float64(len(array))
	// Находит среднее значение
	average = sum / arrLen
	// Возвращает полученное среднее значение
	return average, nil
}

// Sum - суммирует значения из текстового файла
func Sum(fpath string) (sum float64, err error) {
	// Считывает значения из файла
	array, err := readfile.ReadFile(fpath)
	if err != nil {
		return sum, err
	}
	// Через for ... range производит суммирование каждой строчки файла
	for i, _ := range array {
		sum += array[i]
	}
	// Возвращает полученное значение
	return sum, nil
}

// Diff - находи разность между значениями из файла
func Diff(fpath string) (diff float64, err error) {
	// Считывает значения из файла
	array, err := readfile.ReadFile(fpath)
	// Проверка на ошибки
	if err != nil {
		return diff, err
	}
	for i, _ := range array {
		diff -= array[i]
	}
	// Возвращает полученное значение
	return diff, err
}
