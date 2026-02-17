package operation

import (
	"readfile"
)

func Average(fPath string) (average float64, err error) {
	array, err := readfile.ReadFile(fPath)
	if err != nil {
		return average, err
	}
	sum, err := Sum(fPath)
	if err != nil {
		return average, err
	}
	arrLen := float64(len(array))
	average = sum / arrLen
	return average, nil
}
