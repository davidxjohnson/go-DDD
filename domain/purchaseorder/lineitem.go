package purchaseorder

import (
	"go-DDD/domain/product"
)

type lineItem struct {
	product  product.Product
	orderQty int
}

func (l lineItem) New (
	product product.Product,
	orderQty int)(newLineItem lineItem, err error) {

	l.product  = product
	l.orderQty = orderQty

	newLineItem = l

	return
}

