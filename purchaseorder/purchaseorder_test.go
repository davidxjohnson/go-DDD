package purchaseorder

import  (
    "testing"
    "go-DDD/customer"
    "fmt"
)

func TestNewPurchaseOrder(t *testing.T) {
        myPhones := customer.Phones{"work":"212-555-1212","home":"201-555-1212","cell":"908-555-1212"}
        myCustomer, err := customer.NewCustomer( "David", "Johnson", myPhones, "666", "Mockingbird Lane", "Mayberry", "NC", "27030", "US" )
        myPO, err := NewPurchaseOrder( myCustomer )
        // myPO, err := PurchaseOrder.New( myCustomer )
	if err != nil {
		t.Errorf("TestNewPurchaseOrder: NewPurchaseOrder() failed: %s", err.Error() )
	} else {
                fmt.Printf("Successfully create a Purchaseorder: %+v \n", myPO ) 
        }
}
