package main

import (
	"encoding/json"
	"bytes"
	"fmt"
	"os"

	"go-DDD/purchaseorder"

	"github.com/satori/go.uuid"
)

func generateUUID() string {
	return (uuid.Must(uuid.NewV4())).String()
}

//This function takes in a object of type PurchaseOrder for type safety, then encodes it as
//JSON. This provides a []byte which we can then stream/write/etc as needed for persistance
func persistState(po purchaseorder.PurchaseOrder) (string, error) {
	var (
		fileId      = generateUUID()
		lcFile, err = os.Create("/tmp/" + fileId)
		buffer      = new(bytes.Buffer) 
	)

	if err != nil {
		fmt.Println(err)
	}
	
	json.NewEncoder(buffer).Encode(po)
	buffer.WriteTo(lcFile)
	
	return fileId, nil
}

func retrieveState(id string) (purchaseorder.PurchaseOrder, error) {
	po := purchaseorder.PurchaseOrder{}

	return po, nil
}

func main() {
	po := purchaseorder.PurchaseOrder{}

	fmt.Println(persistState(po))
}
