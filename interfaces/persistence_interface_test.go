package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"testing"
	
	"go-DDD/purchaseorder"
	"go-DDD/utility"
)

type dataPersistence struct {
	filename string
}

//This function takes in a object of type PurchaseOrder for type safety, then encodes it as
//JSON. This returns a []byte which we can then stream/write/etc as needed for persistance.
//It returns a string to be used an index id so the object can be retrieved later
func (d dataPersistence)persistState(po purchaseorder.PurchaseOrder) (string, error) {
	var (
		fileID      = utility.Makeuuid()
		lcFile, err = os.Create("/tmp/" + fileID)
		buffer      = new(bytes.Buffer)
	)

	if err != nil {
		return "", err
	}

	json.NewEncoder(buffer).Encode(po)
	buffer.WriteTo(lcFile)

	return fileID, nil
}

func (d dataPersistence) retrieveState(id string) (purchaseorder.PurchaseOrder) {
	po := purchaseorder.PurchaseOrder{}

	return po
}

//This function tests implementation of our interface by creating a variable whose type is 
//the interface we created. After that, we assign a variable of an arbitrary type, to test
//whether that type satisifies the persisted interface
func testInterfaceImplementation(t *testing.T) {
	var (
		persistence Persisted

		dataStore = dataPersistence{filename: "localfile",}
	)

	persistence = dataStore
	fmt.Println(persistence)
}
