package interfaces

import (
	"encoding/json"
	"bytes"
	"os"

	"go-DDD/purchaseorder"

	"github.com/satori/go.uuid"
)

//This interface ensures that any type conforming to the state spec must implement
//a method capable of persisting data and one capable of retrieving data 
type state interface {
	persistState(po purchaseorder.PurchaseOrder) (string, error)
}

func generateUUID() string {
	return (uuid.Must(uuid.NewV4())).String()
}

//This function takes in a object of type PurchaseOrder for type safety, then encodes it as
//JSON. This provides a []byte which we can then stream/write/etc as needed for persistance
//It returns a string which is the index UUID so that the object can be retrieved later
func persistState(po purchaseorder.PurchaseOrder) (string, error) {
	var (
		fileId      = generateUUID()
		lcFile, err = os.Create("/tmp/" + fileId)
		buffer      = new(bytes.Buffer) 
	)

	if err != nil {
		return "", err
	}
	
	json.NewEncoder(buffer).Encode(po)
	buffer.WriteTo(lcFile)
	
	return fileId, nil
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
