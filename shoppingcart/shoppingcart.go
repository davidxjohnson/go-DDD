package shoppingcart

//TODO: remember - don't expose domain objects outside of this current scope 
//TODO: think of this like a transitory state to get us to a persisted purchase order 
import (
	"go-DDD/product"
	"go-DDD/customer"
	"go-DDD/utility"
)

type shoppingcart struct {
	items  map[string][]product.Product
	cartID string 
}

//TODO: replace test string with a valid UUID
func (s *shoppingcart) newShoppingCart() *shoppingcart {
	s.cartID = utility.GenerateUUID()

	return s
}

func (s *shoppingcart) addItem(qty int, cartItem product.Product) {}

//TODO: how do we make sure we target the correct item for deletion?
func (s *shoppingcart) deleteItem(qty int, cartItem product.Product) {}

//JCS - made this a function rather than a method. Passing in a pointer to the obj so we 
//get a reference not a value. This allows us to modify the attribute of the original object
func changeQuantity(qty int, cartItem *product.Product) {
	cartItem.Quantity = quantity
}

func (s *shoppingcart) submitCart(customer customer.Customer) {}
