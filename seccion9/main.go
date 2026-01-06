// A partir de unos precios determinados libre de impuestos y a partir de diferentes tasas de impuestos,
// se pretende calcular para cada unos de los precios, el valor de los nuevos precios luego de anadir los impuestos determinados.

package main

import (
	"fmt"

	"github.com/maxirobledo/curso-go/packages/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.10, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludesPriceJob(taxRate)

		err := priceJob.Process()

		if err != nil {
			fmt.Println("No se pudo procesar la tarea.")
			fmt.Println(err)
		}

	}
}
