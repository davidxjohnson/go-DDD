package shoppingcart

//TODO: remember - don't expose domain objects outside of this current scope 
//TODO: think of this like a transitory state to get us to a persisted purchase order 
import (
	"go-DDD/product"
	"go-DDD/customer"
	"go-DDD/utility"
)

type shoppingcart struct {
	items  map[string][]cartitem
	cartID string 
}

//TODO: replace test string with a valid UUID
func (s *shoppingcart) newShoppingCart() *shoppingcart {
	s.cartID = utility.GenerateUUID()

	return s
}

func (s *shoppingcart) addItem(qty int, item cartitem) {}

//TODO: how do we make sure we target the correct item for deletion?
func (s *shoppingcart) deleteItem(qty int, product product.Product) {}

func (s *shoppingcart) changeQuantity(qty int, product product.Product) {}

func (s *shoppingcart) submitCart(customer customer.Customer) {}
