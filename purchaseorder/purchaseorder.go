package purchaseorder

import (
    "time"
    "go-DDD/customer"
    "go-DDD/utility"
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
func NewPurchaseOrder( aCustomer customer.Customer) (aNewPurchaseOrder PurchaseOrder, err error) {
//func (p PurchaseOrder) New( aCustomer customer.Customer ) ( aNewPurchaseOrder PurchaseOrder, err error ) {
      aNewPurchaseOrder = PurchaseOrder{ orderId: utility.Makeuuid(), orderDate: time.Now(), customer: aCustomer, lineItems: []lineItem{} }
      return
}
