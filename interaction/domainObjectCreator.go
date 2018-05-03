package interaction

import (
	
)
//Struct of structs - create individual data structures for each attribute of the model
//If the specified domain object doesn't match the criteria or is invalid, it's rejected
//The validation is done by the getters and setters methods configured on the object
//Think of the methods in terms of things you will *do* to the object
//Example - find by keys, remove objects, etc. . This would require you to create an obj
//          of type key
//Forget about the infra and focus on the data model 
//Focus on the top-down implementation of the design rather than  from the bottom-uip
//The domain model is how you establish the concept or idea of the object to represent

const maxFieldLength = 100

//Change this to a simple string so it can be validated with a regex
type phonenumber struct {
	areacode int
	prefix   int
	number   int 
}

//Change this to a simple string so it can be validated with a regex 
type email struct {
	domain   string
	username string 
}

type name struct {
	entry string
}

type suffix struct {
	entry string
}

type address struct {
	number  int
	street  string
	city    string
	state   string
	country string
	planet  string
}

type contact struct {
	CellPhone    phonenumber `json:"cell_phone"`
	HomePhone    phonenumber `json:"home_phone"`
	EmailAddress email       `json:"email_address"`
	Firstname    name        `json:"first_name"`
	Lastname     name        `json:"last_name"`
	Suffix       suffix      `json:"suffix"`	
	HomeAddress  address     `json:"home_address"`
	WorkAddress  address     `json:"work_address"`
}

//Rebind these to their datatypes and rather than passing in an argument
//use this binding as a validation mechanism. Focus on their usage as 
//methods rather than as i/o functions (oop paradigms)
func (p phonenumber) validatePhoneNumber(phonenumber interface{}) phonenumber {
	//TODO: regex to extract phone number
	return p
}

func (e email) validateEmailAddress(emailAddress interface{}) email {
	//TODO: regex to validate email address format
	return e
}

func (n name) validateName(contactName interface{}) name {
	return n
}

func (s suffix) validateSuffix(suffix interface{}) suffix { 
	return s
}

func (a address) validateAddress(address interface{}) address {
	return a
}

func createContact(contactFields ...interface{}) {
	//TODO: Figure out a way to validate inbound, ambiguously typed data
	//TODO: so that we can call the appropriate function for that data
	//TODO: What about goroutines? Each method is operated on its 
	//TODO: own thread
	//What if we bind each field in the struct to the appropriate
	//constructor for that field and then loop those fields with the 
	//passed-in parameter from the function call??
	
	for _, contact := range contactFields {

	}
}