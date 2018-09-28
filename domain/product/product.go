package product

import (
	"go-DDD/utility"
)

type Product struct {
	id            string
	shortName     string
	description   string
	price         float64
}

func (p Product) New(
	shortName   string,
	description string,
	price       float64) (newProduct Product, err error) {

	p.id          = utility.Makeuuid()	
	p.shortName   = shortName
	p.description = description
	p.price       = price

	newProduct = p

	return 
}
