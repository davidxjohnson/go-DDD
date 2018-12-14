package interfaces

import (
	"fmt"
	"go-DDD/customer"
	"go-DDD/purchaseorder"
	"mapstore"
	"testing"
)

type dataPersistence struct {
	table mapstore.Table
}

//This function takes in a object of type PurchaseOrder for type safety, then encodes it as
//JSON. This returns a []byte which we can then stream/write/etc as needed for persistance.
//It returns a string to be used an index id so the object can be retrieved later
func (d dataPersistence) persistState(po purchaseorder.PurchaseOrder) (id string, err error) {
	var data interface{} = po
	var myRow = mapstore.Row{"po": data}
	id, _ = d.table.AddRow(myRow)
	d.table.CommitTable()
	return
}

func (d dataPersistence) retrieveState(id string) (po purchaseorder.PurchaseOrder) {
	myRow, _ := d.table.FindRowByID(id)
	po = myRow["po"].(purchaseorder.PurchaseOrder)
	return
}

//This function tests implementation of our interface by creating a variable whose type is
//the interface we created. After that, we assign a variable of an arbitrary type, to test
//whether that type satisifies the persisted interface
func TestInterfaceImplementation(t *testing.T) {
	table, _ := mapstore.NewTable("./testdata/testdata.json")
	var store Persisted = dataPersistence{table: table}
	myPhones := customer.Phones{"work": "212-555-1212", "home": "201-555-1212", "cell": "908-555-1212"}
	myCustomer, _ := customer.Customer{}.New("David", "Johnson", myPhones, "666", "Mockingbird Lane", "Mayberry", "NC", "27030", "US")
	myPO1, _ := purchaseorder.PurchaseOrder{}.New(myCustomer)
	recordKey, _ := store.persistState(myPO1)
	fmt.Printf("Stored record key = %s, store content = %++v \n", recordKey, store)
	myPO2 := store.retrieveState(recordKey)
	fmt.Printf("Retrieved po = %++v \n", myPO2)
}
