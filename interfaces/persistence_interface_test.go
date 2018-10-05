package interfaces

import (
	"bytes"
	"encoding/json"
	"os"

	"go-DDD/purchaseorder"
	"go-DDD/utility"
)

// implement interface here ... secret sauce
// #####

//This function takes in a object of type PurchaseOrder for type safety, then encodes it as
//JSON. This returns a []byte which we can then stream/write/etc as needed for persistance.
//It returns a string to be used an index id so the object can be retrieved later
func persistState(po purchaseorder.PurchaseOrder) (string, error) {
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

func retrieveState(id string) (purchaseorder.PurchaseOrder, error) {
	po := purchaseorder.PurchaseOrder{}

	return po, nil
}

//func main() {
//	po := purchaseorder.PurchaseOrder{}
//
//	fmt.Println(persistState(po))
//}
