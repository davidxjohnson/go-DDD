package purchaseorder

import (
	"fmt"
	"testing"

	"go-DDD/product"
)

func TestNewCustomer(t *testing.T) {
	product,_ := product.Product{}.New(
		"flux capacitor", "permits time travel", 100.00)

	myLineItem, err := lineItem{}.New(product, 1) 

	if err != nil {
		t.Errorf("TestNewLineItem: NewLineItem() failed: %s", err.Error())
	} else {
		fmt.Printf("Successfully create a lineItem: %+v \n", myLineItem)
	}
}
