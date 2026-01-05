package conversion

import (
	"errors"
	"strconv"
)

func StringToFloat(str []string) ([]float64, error) {
	var floats []float64

	for _, strVal := range str {
		floatVal, err := strconv.ParseFloat(strVal, 64)

		if err != nil {
			return nil, errors.New("Error al convertir string a float")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}
