package interfaces

import (
	"go-DDD/purchaseorder"
)

//This interface ensures that any type conforming to the state spec must implement
//a method capable of persisting data and one capable of retrieving data
type Persisted interface {
	persistState(po purchaseorder.PurchaseOrder) (string, error)
	retrieveState(id string) purchaseorder.PurchaseOrder
}
