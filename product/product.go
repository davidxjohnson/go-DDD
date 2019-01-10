package product

//JCS - Added a quantity here, since the number of products should be an attribute of that 
//product. Made it public so it can be modified externally via a function in shopping cart
type Product struct {
	id            string
	shortName     string
	description   string
	price         float64
	Quantity      int
}
