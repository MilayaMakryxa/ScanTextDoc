// Данная программа считывает данные из файла
package readfile

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(fPath string) (data [3]float64, err error) {
	// Открываем файл с данными
	file, err := os.Open(fPath)
	// Проверка на ошибки открытия файла
	if err != nil {
		return data, err
	}
	// Создаем новый сканер для file
	scan := bufio.NewScanner(file)
	// Считываем данные с файла
	for i := 0; scan.Scan(); i++ {
		// Загоняет данные в массив и преобразует строки в float64
		data[i], err = strconv.ParseFloat(scan.Text(), 64)
	}
	// Проверка на оишбки при сканировании
	if scan.Err() != nil {
		return data, scan.Err()
	}
	// Проверка на ошибки при закрытии файла
	err = file.Close()
	if err != nil {
		return data, err
	}
	return data, nil
}
