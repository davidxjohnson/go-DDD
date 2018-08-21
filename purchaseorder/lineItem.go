package purchaseorder

import (
	"go-DDD/product"
)

type lineItem struct {
	product  product.Product
	orderQty int
}
