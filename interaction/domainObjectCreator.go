package interaction

import (
	"strings"
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

type phonenumber struct {
	areacode int
	prefix   int
	number   int 
}

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

func validatePhoneNumber(number phonenumber) bool {
	return false
}

func validateEmailAddress(emailAddress email) bool {
	if(len(emailAddress.username) + len(emailAddress.domain)) <= maxFieldLength && strings.Contains(e.domain, "@"){
		return true
	}	
	return false
}

func (n name) validateName(contactName name) bool {
	return false	
}

func (s suffix) validateSuffix(suffixEntry suffix) bool {
	return false
}

func (a address) validateAddress(addressEntry address) bool {
	return false
}

func (c contact) createContact(contactFields ...interface{}) {
	for _, contactField := range contactFields {
		switch contactField.(type) {
		case phonenumber:
			assertedPhoneNumber := contactField.(phonenumber)
			
			if validatePhoneNumber(assertedPhoneNumber) == true {
				c.HomePhone = assertedPhoneNumber
			}
		}
	}
}