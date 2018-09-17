package product

import (
	"fmt"
	"testing"
)

func TestNewProduct(t *testing.T) {
	myProduct, err := Product{}.New(
		"plutonium",
		"used for powering a flux capacitor for purposes of time travel",
		5.0)
	
	if err != nil {
		t.Errorf("TestNewCustomer: NewCustomer() failed: %s", err.Error())
	} else {
		fmt.Printf("Successfully created a product: %+v \n", myProduct)
	}
}
