package product

//JCS - Added a quantity here, since the number of products should be an attribute of that 
//product. Made it public so it can be modified externally via a function in shopping cart
type Product struct {
	id            string
	shortName     string
	description   string
	price         float64
	Qty           int
}

//TODO: should we use a function for an init or should we bind it as a method to a datatype?
//TODO: my thought is shouldn't we instantiate an object using a function and then call 
//TODO: our methods on the object that it returns?
func (p *Product) NewProduct(id, shortName, description string, price float64, qty int) {
	p.id          = id
	p.shortName   = shortName
	p.description = description
	p.price       = price
	p.Qty         = qty
}
