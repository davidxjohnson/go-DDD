package shoppingcart

//TODO: rework this to remove customer, add an ID, and a mechanism to create a unique cart
//TODO: for each customer
//TODO: remember - don't expose domain objects outside of this current scope 
//TODO: think of this like a transitory state to get us to a persisted purchase order 
import (
	"go-DDD/product"
	"go-DDD/customer"
)

type shoppingcart struct {
	items  []cartitem
	cartID string 
}

type cartitem struct {
	Product  product.Product
	Quantity int
}

//TODO: replace test string with a valid UUID
func (s *shoppingcart) newShoppingCart() *shoppingcart {
	s.cartID = "test"

	return s
}

func (s *shoppingcart) addItem(qty int, item cartitem) {
	if s.items == nil {
		s.items = []cartitem{item,}
	}

	s.items = append(s.items, item)
}

func (s *shoppingcart) deleteItem(qty int, product product.Product) {
}

func (s *shoppingcart) changeQuantity(qty int, product product.Product) {
}

func (s *shoppingcart) submitCart(customer customer.Customer) {
}
