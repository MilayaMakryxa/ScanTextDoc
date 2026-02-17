package operation

import "readfile"

func Sum(fpath string) (sum float64, err error) {
	array, err := readfile.ReadFile(fpath)
	if err != nil {
		return sum, err
	}
	for i, _ := range array {
		sum += array[i]
	}
	return sum, nil
}
