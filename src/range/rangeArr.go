package rangeArr

import "fmt"

func Range() (fNum, sNum int, err error) {
	fmt.Println("Введите сколько строк файла необходимо отобразить.")
	fmt.Print("Первая строка файла:")
	fmt.Scan(&fNum)
	fmt.Print("Последняя строка файла:")
	fmt.Scan(&sNum)
	return fNum, sNum, err

}
