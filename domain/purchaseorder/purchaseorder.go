package purchaseorder

import (
	"go-DDD/domain/customer"
	"go-DDD/utility"
	"time"
)

type PurchaseOrder struct {
	orderId      string
	orderDate    time.Time
	shippingDate time.Time
	customer     customer.Customer
	lineItems    []lineItem
}

// type methods

// New returns a newly initialized purchase order
func (p PurchaseOrder) New(aCustomer customer.Customer) (newPO PurchaseOrder, err error) {
	p.orderId = utility.Makeuuid()
	p.orderDate = time.Now()
	p.customer = aCustomer
	p.lineItems = []lineItem{}
	newPO = p
	return
}
