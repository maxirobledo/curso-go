package products

import (
	"errors"
	"fmt"
	"time"
)

type Product struct {
	Title      string    "json:\"title\""
	Id         int       "json:\"id\""
	Price      float64   "json:\"price\""
	Created_at time.Time "json:\"created_at\""
}

func New(title string, id int, price float64) (*Product, error) {
	if title == "" || id == 0 || price <= 0 {
		return nil, errors.New("ERROR - Se ingresaron valores vacios o no permitidos")
	}

	return &Product{
		Title:      title,
		Id:         id,
		Price:      price,
		Created_at: time.Now(),
	}, nil

}

func GetProduct(p *Product) string {
	return fmt.Sprintf("id: %d\nTitle: %s\nPrice: %2.f", p.Id, p.Title, p.Price)
}
