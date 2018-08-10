package purchaseorder

type PurchaseOrder struct {
	ordernumber  int 
	shippingdate string
	city         string
	street       string
	zipCode      int 
	ordereditem  orderedItem
	quantity     int
	product      productDescription
}


