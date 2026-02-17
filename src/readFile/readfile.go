// Данная программа считывает данные из файла
package readfile

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFile(fPath string) (data []float64, err error) {
	// Открываем файл с данными
	file, err := os.Open(fPath)
	// Проверка на ошибки открытия файла
	if err != nil {
		return data, err
	}
	// Создаем новый сканер для file
	scan := bufio.NewScanner(file)
	// Считываем данные с файла
	for scan.Scan() {
		// Полученные данные сохраняет в временную переменную
		number, err := strconv.ParseFloat(scan.Text(), 64)
		// Проверка на оишбки при заполнении
		if err != nil {
			return data, err
		}
		// Добавляет к сегменту новые данные
		data = append(data, number)
	}
	// Проверка на ошибки при сканировании
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
