package customer

import  (
    "testing"
    "fmt"
)

func TestNewCustomer(t *testing.T) {
        myPhones := Phones{"work":"212-555-1212","home":"201-555-1212","cell":"908-555-1212"}
        myCustomer, err := NewCustomer( "David", "Johnson", myPhones, "666", "Mockingbird Lane", "Mayberry", "NC", "27030", "US" )
        if err != nil {
                t.Errorf("TestNewCustomer: NewCustomer() failed: %s", err.Error() )
        } else {
                fmt.Printf("Successfully create a Customer: %+v \n", myCustomer )
        }
}
