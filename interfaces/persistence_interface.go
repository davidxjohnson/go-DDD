package interfaces

import (
	"go-DDD/purchaseorder"
)

// Presisted ... storage interface for storing and retrieving PurchaseOrder
type Persisted interface {
	persistState(po purchaseorder.PurchaseOrder) (string, error)
	retrieveState(id string) purchaseorder.PurchaseOrder
}
