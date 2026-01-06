package prices

import (
	"fmt"

	"github.com/maxirobledo/curso-go/packages/conversion"
	"github.com/maxirobledo/curso-go/packages/fileops"
)

type TaxIncludesPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludesPrices map[string]string
}

func NewTaxIncludesPriceJob(taxRate float64) *TaxIncludesPriceJob {
	return &TaxIncludesPriceJob{
		TaxRate: taxRate,
	}
}

func (job *TaxIncludesPriceJob) Process() error {

	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludesPrice := price * (1 + job.TaxRate)

		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludesPrice)
	}

	job.TaxIncludesPrices = result

	fileops.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

	return nil
}

func (job *TaxIncludesPriceJob) LoadData() error {
	lines, err := fileops.ReadLines("prices.txt")

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloat(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices
	return nil
}
