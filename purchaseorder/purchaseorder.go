package purchaseorder

import (
	"go-DDD/customer"
	"go-DDD/utility"
	"time"
)

type PurchaseOrder struct {
	orderId      string            `json:"orderId"`
	orderDate    time.Time         `json:"orderDate"`
	shippingDate time.Time         `json:"shippingDate"`
	customer     customer.Customer `json:"customer"`
	lineItems    []lineItem        `json:"lineItems"`
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
