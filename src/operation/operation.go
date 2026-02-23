package operation

// Average - вычисляет среднее значение из тектового файла
func Average(data []float64) (average float64, err error) {
	// Суммирует полученные значения из файла функции Sum
	sum, err := Sum(data)
	// Проверка на ошибки при расчете
	if err != nil {
		return sum, err
	}
	// Считает длину массива
	lenArr := float64(len(data))
	// Считает среднее
	average = sum / lenArr
	return average, nil
}

// Sum - суммирует значения из текстового файла
func Sum(data []float64) (sum float64, err error) {
	// Через for ... range производит суммирование каждой строчки файла
	for i, _ := range data {
		sum += data[i]
	}
	// Возвращает полученное значение
	return sum, nil
}

// Diff - находи разность между значениями из файла
func Diff(data []float64) (diff float64, err error) {
	for i, _ := range data {
		diff -= data[i]
	}
	// Возвращает полученное значение
	return diff, err
}
